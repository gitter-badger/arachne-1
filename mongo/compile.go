package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/bmeg/arachne/aql"
	"github.com/bmeg/arachne/engine/core"
	"github.com/bmeg/arachne/gdbi"
	"github.com/bmeg/arachne/jsonpath"
	"github.com/bmeg/arachne/protoutil"
	"gopkg.in/mgo.v2/bson"
)

type MongoCompiler struct {
	db       *Graph
	compiler gdbi.Compiler
}

func NewCompiler(db *Graph) gdbi.Compiler {
	return &MongoCompiler{db: db, compiler: core.NewCompiler(db)}
}

type MongoPipeline struct {
	db              *Graph
	startCollection string
	query           []bson.M
	dataType        gdbi.DataType
	markTypes       map[string]gdbi.DataType
}

func (comp *MongoCompiler) Compile(stmts []*aql.GraphStatement) (gdbi.Pipeline, error) {
	query := []bson.M{}
	startCollection := ""
	lastType := gdbi.NoData
	markTypes := map[string]gdbi.DataType{}

	vertCol := fmt.Sprintf("%s_vertices", comp.db.graph)
	edgeCol := fmt.Sprintf("%s_edges", comp.db.graph)

	for _, gs := range stmts {
		switch stmt := gs.GetStatement().(type) {
		case *aql.GraphStatement_V:
			if lastType != gdbi.NoData {
				return &MongoPipeline{}, fmt.Errorf(`"V" statement is only valid at the beginning of the traversal`)
			}
			startCollection = vertCol
			ids := protoutil.AsStringList(stmt.V)
			if len(ids) > 0 {
				query = append(query, bson.M{"$match": bson.M{"_id": bson.M{"$in": ids}}})
			}
			lastType = gdbi.VertexData

		case *aql.GraphStatement_E:
			if lastType != gdbi.NoData {
				return &MongoPipeline{}, fmt.Errorf(`"E" statement is only valid at the beginning of the traversal`)
			}
			startCollection = edgeCol
			ids := protoutil.AsStringList(stmt.E)
			if len(ids) > 0 {
				query = append(query, bson.M{"$match": bson.M{"_id": bson.M{"$in": ids}}})
			}
			lastType = gdbi.EdgeData

		case *aql.GraphStatement_In:
			labels := protoutil.AsStringList(stmt.In)
			if lastType == gdbi.VertexData {
				query = append(query,
					bson.M{
						"$lookup": bson.M{
							"from": edgeCol,
							"let":  bson.M{"vid": "$_id", "marks": "$marks"},
							"pipeline": []bson.M{
								{
									"$match": bson.M{
										"$expr": bson.M{
											"$eq": []string{"$to", "$$vid"},
										},
									},
								},
								{
									"$addFields": bson.M{"marks": "$$marks"},
								},
							},
							"as": "dst",
						},
					},
				)
				query = append(query, bson.M{"$unwind": "$dst"})
				query = append(query, bson.M{"$project": bson.M{"_id": "$dst._id", "label": "$dst.label", "data": "$dst.data", "to": "$dst.to", "from": "$dst.from", "marks": "$dst.marks"}})
			}
			if len(labels) > 0 {
				query = append(query, bson.M{"$match": bson.M{"label": bson.M{"$in": labels}}})
			}
			query = append(query,
				bson.M{
					"$lookup": bson.M{
						"from": vertCol,
						"let":  bson.M{"from": "$from", "marks": "$marks"},
						"pipeline": []bson.M{
							{
								"$match": bson.M{
									"$expr": bson.M{
										"$eq": []string{"$_id", "$$from"},
									},
								},
							},
							{
								"$addFields": bson.M{"marks": "$$marks"},
							},
						},
						"as": "dst",
					},
				},
			)
			query = append(query, bson.M{"$unwind": "$dst"})
			query = append(query, bson.M{"$project": bson.M{"_id": "$dst._id", "label": "$dst.label", "data": "$dst.data", "marks": "$dst.marks"}})
			lastType = gdbi.VertexData

		case *aql.GraphStatement_Out:
			labels := protoutil.AsStringList(stmt.Out)
			if lastType == gdbi.VertexData {
				query = append(query,
					bson.M{
						"$lookup": bson.M{
							"from": edgeCol,
							"let":  bson.M{"vid": "$_id", "marks": "$marks"},
							"pipeline": []bson.M{
								{
									"$match": bson.M{
										"$expr": bson.M{
											"$eq": []string{"$from", "$$vid"},
										},
									},
								},
								{
									"$addFields": bson.M{"marks": "$$marks"},
								},
							},
							"as": "dst",
						},
					},
				)
				query = append(query, bson.M{"$unwind": "$dst"})
				query = append(query, bson.M{"$project": bson.M{"_id": "$dst._id", "label": "$dst.label", "data": "$dst.data", "to": "$dst.to", "from": "$dst.from", "marks": "$dst.marks"}})
			}
			if len(labels) > 0 {
				query = append(query, bson.M{"$match": bson.M{"label": bson.M{"$in": labels}}})
			}
			query = append(query,
				bson.M{
					"$lookup": bson.M{
						"from": vertCol,
						"let":  bson.M{"to": "$to", "marks": "$marks"},
						"pipeline": []bson.M{
							{
								"$match": bson.M{
									"$expr": bson.M{
										"$eq": []string{"$_id", "$$to"},
									},
								},
							},
							{
								"$addFields": bson.M{"marks": "$$marks"},
							},
						},
						"as": "dst",
					},
				},
			)
			query = append(query, bson.M{"$unwind": "$dst"})
			query = append(query, bson.M{"$project": bson.M{"_id": "$dst._id", "label": "$dst.label", "data": "$dst.data", "marks": "$dst.marks"}})
			lastType = gdbi.VertexData

		case *aql.GraphStatement_Both:
			labels := protoutil.AsStringList(stmt.Both)
			if lastType == gdbi.VertexData {
				query = append(query,
					bson.M{
						"$lookup": bson.M{
							"from": edgeCol,
							"let":  bson.M{"vid": "$_id", "marks": "$marks"},
							"pipeline": []bson.M{
								{
									"$match": bson.M{
										"$expr": bson.M{
											"$or": []bson.M{
												{"$eq": []string{"$to", "$$vid"}},
												{"$eq": []string{"$from", "$$vid"}},
											},
										},
									},
								},
								{"$addFields": bson.M{"vid": "$$vid", "marks": "$$marks"}},
							},
							"as": "dst",
						},
					},
				)
				query = append(query, bson.M{"$unwind": "$dst"})
				query = append(query, bson.M{"$project": bson.M{"_id": "$dst._id", "label": "$dst.label", "data": "$dst.data", "to": "$dst.to", "from": "$dst.from", "marks": "$dst.marks", "vid": "$dst.vid"}})
			}
			if len(labels) > 0 {
				query = append(query, bson.M{"$match": bson.M{"label": bson.M{"$in": labels}}})
			}
			query = append(query,
				bson.M{
					"$lookup": bson.M{
						"from": vertCol,
						"let":  bson.M{"to": "$to", "from": "$from", "marks": "$marks", "vid": "$vid"},
						"pipeline": []bson.M{
							{
								"$match": bson.M{
									"$expr": bson.M{
										"$and": []bson.M{
											{
												"$or": []bson.M{
													{"$eq": []string{"$_id", "$$from"}},
													{"$eq": []string{"$_id", "$$to"}},
												},
											},
											{
												"$ne": []string{"$_id", "$$vid"},
											},
										},
									},
								},
							},
							{
								"$addFields": bson.M{"marks": "$$marks"},
							},
						},
						"as": "dst",
					},
				},
			)
			query = append(query, bson.M{"$unwind": "$dst"})
			query = append(query, bson.M{"$project": bson.M{"_id": "$dst._id", "label": "$dst.label", "data": "$dst.data", "marks": "$dst.marks"}})
			lastType = gdbi.VertexData

		case *aql.GraphStatement_InEdge:
			if lastType != gdbi.VertexData {
				return &MongoPipeline{}, fmt.Errorf(`"inEdge" statement is only valid for the vertex type not: %s`, lastType.String())
			}
			query = append(query,
				bson.M{
					"$lookup": bson.M{
						"from": edgeCol,
						"let":  bson.M{"vid": "$_id", "marks": "$marks"},
						"pipeline": []bson.M{
							{
								"$match": bson.M{
									"$expr": bson.M{
										"$eq": []string{"$to", "$$vid"},
									},
								},
							},
							{
								"$addFields": bson.M{"marks": "$$marks"},
							},
						},
						"as": "dst",
					},
				},
			)
			query = append(query, bson.M{"$unwind": "$dst"})
			query = append(query, bson.M{"$project": bson.M{"_id": "$dst._id", "label": "$dst.label", "data": "$dst.data", "to": "$dst.to", "from": "$dst.from", "marks": "$dst.marks"}})
			labels := protoutil.AsStringList(stmt.InEdge)
			if len(labels) > 0 {
				query = append(query, bson.M{"$match": bson.M{"label": bson.M{"$in": labels}}})
			}
			lastType = gdbi.EdgeData

		case *aql.GraphStatement_OutEdge:
			if lastType != gdbi.VertexData {
				return &MongoPipeline{}, fmt.Errorf(`"outEdge" statement is only valid for the vertex type not: %s`, lastType.String())
			}
			query = append(query,
				bson.M{
					"$lookup": bson.M{
						"from": edgeCol,
						"let":  bson.M{"vid": "$_id", "marks": "$marks"},
						"pipeline": []bson.M{
							{
								"$match": bson.M{
									"$expr": bson.M{
										"$eq": []string{"$from", "$$vid"},
									},
								},
							},
							{
								"$addFields": bson.M{"marks": "$$marks"},
							},
						},
						"as": "dst",
					},
				},
			)
			query = append(query, bson.M{"$unwind": "$dst"})
			query = append(query, bson.M{"$project": bson.M{"_id": "$dst._id", "label": "$dst.label", "data": "$dst.data", "to": "$dst.to", "from": "$dst.from", "marks": "$dst.marks"}})
			labels := protoutil.AsStringList(stmt.OutEdge)
			if len(labels) > 0 {
				query = append(query, bson.M{"$match": bson.M{"label": bson.M{"$in": labels}}})
			}
			lastType = gdbi.EdgeData

		case *aql.GraphStatement_BothEdge:
			if lastType != gdbi.VertexData {
				return &MongoPipeline{}, fmt.Errorf(`"bothEdge" statement is only valid for the vertex type not: %s`, lastType.String())
			}
			query = append(query,
				bson.M{
					"$lookup": bson.M{
						"from": edgeCol,
						"let":  bson.M{"vid": "$_id", "marks": "$marks"},
						"pipeline": []bson.M{
							{
								"$match": bson.M{
									"$expr": bson.M{
										"$or": []bson.M{
											{"$eq": []string{"$to", "$$vid"}},
											{"$eq": []string{"$from", "$$vid"}},
										},
									},
								},
							},
							{
								"$addFields": bson.M{"marks": "$$marks"},
							},
						},
						"as": "dst",
					},
				},
			)
			query = append(query, bson.M{"$unwind": "$dst"})
			query = append(query, bson.M{"$project": bson.M{"_id": "$dst._id", "label": "$dst.label", "data": "$dst.data", "to": "$dst.to", "from": "$dst.from", "marks": "$dst.marks"}})
			labels := protoutil.AsStringList(stmt.BothEdge)
			if len(labels) > 0 {
				query = append(query, bson.M{"$match": bson.M{"label": bson.M{"$in": labels}}})
			}
			lastType = gdbi.EdgeData

		case *aql.GraphStatement_Where:
			if lastType != gdbi.VertexData && lastType != gdbi.EdgeData {
				return &MongoPipeline{}, fmt.Errorf(`"distinct" statement is only valid for edge or vertex types not: %s`, lastType.String())
			}
			// TODO

		case *aql.GraphStatement_Limit:
			query = append(query,
				bson.M{"$limit": stmt.Limit})

		case *aql.GraphStatement_Offset:
			query = append(query,
				bson.M{"$skip": stmt.Offset})

		case *aql.GraphStatement_Count:
			query = append(query,
				bson.M{"$count": "count"})
			lastType = gdbi.CountData

		case *aql.GraphStatement_Distinct:
			if lastType != gdbi.VertexData && lastType != gdbi.EdgeData {
				return &MongoPipeline{}, fmt.Errorf(`"distinct" statement is only valid for edge or vertex types not: %s`, lastType.String())
			}
			//TODO

		case *aql.GraphStatement_Mark:
			if lastType == gdbi.NoData {
				return &MongoPipeline{}, fmt.Errorf(`"mark" statement is not valid at the beginning of a traversal`)
			}
			if stmt.Mark == "" {
				return &MongoPipeline{}, fmt.Errorf(`"mark" statement cannot have an empty name`)
			}
			if err := aql.ValidateFieldName(stmt.Mark); err != nil {
				return &MongoPipeline{}, fmt.Errorf(`"mark" statement invalid; %v`, err)
			}
			if stmt.Mark == jsonpath.Current {
				return &MongoPipeline{}, fmt.Errorf(`"mark" statement invalid; uses reserved name %s`, jsonpath.Current)
			}
			markTypes[stmt.Mark] = lastType
			query = append(query, bson.M{"$addFields": bson.M{"marks": bson.M{stmt.Mark: "$_id"}}})

		case *aql.GraphStatement_Select:
			if lastType != gdbi.VertexData && lastType != gdbi.EdgeData {
				return &MongoPipeline{}, fmt.Errorf(`"select" statement is only valid for edge or vertex types not: %s`, lastType.String())
			}
			if len(stmt.Select.Marks) == 0 {
				return &MongoPipeline{}, fmt.Errorf(`"select" statement has an empty list of mark names`)
			}
			selection := bson.M{}
			for _, mark := range stmt.Select.Marks {
				selection["marks."+mark] = 1
			}
			query = append(query, bson.M{"$project": selection})
			lastType = gdbi.SelectionData

		case *aql.GraphStatement_Render:
			if lastType != gdbi.VertexData && lastType != gdbi.EdgeData {
				return &MongoPipeline{}, fmt.Errorf(`"render" statement is only valid for edge or vertex types not: %s`, lastType.String())
			}
			//TODO
			lastType = gdbi.RenderData

		case *aql.GraphStatement_Fields:
			if lastType != gdbi.VertexData && lastType != gdbi.EdgeData {
				return &MongoPipeline{}, fmt.Errorf(`"fields" statement is only valid for edge or vertex types not: %s`, lastType.String())
			}
			//TODO

		case *aql.GraphStatement_Aggregate:
			if lastType != gdbi.VertexData {
				return &MongoPipeline{}, fmt.Errorf(`"aggregate" statement is only valid for vertex types not: %s`, lastType.String())
			}
			//TODO
			lastType = gdbi.AggregationData

		default:
			return &MongoPipeline{}, fmt.Errorf("unknown statement type")
		}
	}
	return &MongoPipeline{comp.db, startCollection, query, lastType, markTypes}, nil
}

func (pipe *MongoPipeline) DataType() gdbi.DataType {
	return pipe.dataType
}

func (pipe *MongoPipeline) MarkTypes() map[string]gdbi.DataType {
	return pipe.markTypes
}

func (pipe *MongoPipeline) Processors() []gdbi.Processor {
	return []gdbi.Processor{&MongoProcessor{pipe.db, pipe.startCollection, pipe.query, pipe.dataType, pipe.markTypes}}
}

type MongoProcessor struct {
	db              *Graph
	startCollection string
	query           []bson.M
	dataType        gdbi.DataType
	markTypes       map[string]gdbi.DataType
}

func getDataElement(data map[string]interface{}) *gdbi.DataElement {
	d := &gdbi.DataElement{}
	if x, ok := data["_id"]; ok {
		d.ID = x.(string)
	}
	if x, ok := data["label"]; ok {
		d.Label = x.(string)
	}
	if x, ok := data["data"]; ok {
		d.Data = x.(map[string]interface{})
	}
	if x, ok := data["to"]; ok {
		d.To = x.(string)
	}
	if x, ok := data["from"]; ok {
		d.From = x.(string)
	}
	return d
}

func (proc *MongoProcessor) Process(ctx context.Context, man gdbi.Manager, in gdbi.InPipe, out gdbi.OutPipe) context.Context {
	log.Printf("Running Mongo Processor: %+v", proc.query)

	go func() {
		session := proc.db.ar.pool.Get()
		defer proc.db.ar.pool.Put(session)
		defer close(out)

		initCol := session.DB(proc.db.ar.database).C(proc.startCollection)
		for t := range in {
			iter := initCol.Pipe(proc.query).Iter()
			result := map[string]interface{}{}
			for iter.Next(&result) {
				log.Printf("MongoPipeline result: %+v", result)
				select {
				case <-ctx.Done():
					return
				default:
				}

				switch proc.dataType {
				case gdbi.CountData:
					eo := &gdbi.Traveler{}
					if x, ok := result["count"]; ok {
						eo.Count = uint32(x.(int))
					}
					out <- eo

				case gdbi.RenderData:
					log.Println("MongoProcessor for gdbi.RenderData not implemented")
					//TODO

				case gdbi.SelectionData:
					selections := map[string]*gdbi.DataElement{}
					if marks, ok := result["marks"]; ok {
						if marks, ok := marks.(map[string]interface{}); ok {
							for k, v := range marks {
								gid, ok := v.(string)
								if !ok {
									log.Println("Failed to process selection data: %+v", v)
								}
								de := &gdbi.DataElement{}
								switch proc.markTypes[k] {
								case gdbi.VertexData:
									v := proc.db.GetVertex(gid, true)
									if v != nil {
										de = &gdbi.DataElement{
											ID:    v.Gid,
											Label: v.Label,
											Data:  v.GetDataMap(),
										}
									}
								case gdbi.EdgeData:
									e := proc.db.GetEdge(gid, true)
									if e != nil {
										de = &gdbi.DataElement{
											ID:    e.Gid,
											Label: e.Label,
											From:  e.From,
											To:    e.To,
											Data:  e.GetDataMap(),
										}
									}
								}
								selections[k] = de
							}
						}
					}
					out <- &gdbi.Traveler{Selections: selections}

				case gdbi.AggregationData:
					log.Println("MongoProcessor for gdbi.AggregationData not implemented")
					//TODO

				default:
					out <- t.AddCurrent(getDataElement(result))
				}
			}
			if err := iter.Err(); err != nil {
				log.Println("Mongo traversal error:", err)
				continue
			}
		}
	}()

	return ctx
}
