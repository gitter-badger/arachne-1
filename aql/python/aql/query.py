from __future__ import absolute_import, print_function, unicode_literals

import json
import sys
import urllib2

from aql.util import do_request


class Query:
    def __init__(self, url):
        self.query = []
        self.url = url

    def __append(self, part):
        q = self.__class__(self.url)
        q.query = self.query[:]
        q.query.append(part)
        return q

    def V(self, id=[]):
        """
        Start the query at a vertex.

        "id" is an ID or a list of vertex IDs to start from. Optional.
        """
        if not isinstance(id, list):
            id = [id]
        return self.__append({"v": id})

    def E(self, id=[]):
        """
        Start the query at an edge.

        "id" is an ID to start from. Optional.
        """
        if not isinstance(id, list):
            id = [id]
        return self.__append({"e": id})

    def where(self, expression):
        """
        Filter vertex/edge based on properties.
        """
        return self.__append({"where": expression})

    def fields(self, fields=[]):
        """
        Select document properties to be returned in document.
        """
        if not isinstance(fields, list):
            fields = [fields]
        return self.__append({"fields": fields})

    def in_(self, label=[]):
        """
        Follow an incoming edge to the source vertex.

        "label" is the label of the edge to follow.
        "label" can be a list.
        """
        if not isinstance(label, list):
            label = [label]
        return self.__append({"in": label})

    def out(self, label=[]):
        """
        Follow an outgoing edge to the destination vertex.

        "label" is the label of the edge to follow.
        "label" can be a list.
        """
        if not isinstance(label, list):
            label = [label]
        return self.__append({"out": label})

    def both(self, label=[]):
        """
        Follow both incoming and outgoing edges to vertices.

        "label" is the label of the edge to follow.
        "label" can be a list.
        """
        if not isinstance(label, list):
            label = [label]
        return self.__append({"both": label})

    def inEdge(self, label=[]):
        """
        Move from a vertex to an incoming edge.

        "label" is the label of the edge to move to.
        "label" can be a list.

        Must be called from a vertex.
        """
        if not isinstance(label, list):
            label = [label]
        return self.__append({"in_edge": label})

    def outEdge(self, label=[]):
        """
        Move from a vertex to an outgoing edge.

        "label" is the label of the edge to move to.
        "label" can be a list.

        Must be called from a vertex.
        """
        if not isinstance(label, list):
            label = [label]
        return self.__append({"out_edge": label})

    def bothEdge(self, label=[]):
        """
        Move from a vertex to incoming/outgoing edges.

        "label" is the label of the edge to move to.
        "label" can be a list.

        Must be called from a vertex.
        """
        if not isinstance(label, list):
            label = [label]
        return self.__append({"both_edge": label})

    def mark(self, name):
        """
        Mark the current vertex/edge with the given name.

        Used to return elements from select().
        """
        return self.__append({"mark": name})

    def select(self, marks):
        """
        Returns rows of marked elements, with one item for each mark.

        "marks" is a list of mark names.
        The rows returned are all combinations of marks, e.g.
        [
            [A1, B1],
            [A1, B2],
            [A2, B1],
            [A2, B2],
        ]
        """
        if not isinstance(marks, list):
            marks = [marks]
        return self.__append({"select": {"marks": marks}})

    def limit(self, n):
        """
        Limits the number of results returned.
        """
        return self.__append({"limit": n})

    def offset(self, n):
        """
        Offset the results returned.
        """
        return self.__append({"offset": n})

    def count(self):
        """
        Return the number of results, instead of the elements.
        """
        return self.__append({"count": ""})

    def distinct(self, props=[]):
        """
        Select distinct elements based on the provided property list.
        """
        if not isinstance(props, list):
            props = [props]
        return self.__append({"distinct": props})

    def match(self, queries):
        """
        Intersect multiple queries.
        """
        if not isinstance(queries, list):
            raise TypeError("match expects an array")
        if not all(isinstance(i, Query) for i in queries):
            raise TypeError("expected all aruments to match to be a \
            Query instance")
        mq = []
        for i in queries:
            mq.append({"query": i.query})
        return self.__append({"match": {"queries": mq}})

    def render(self, template):
        """
        Render output of query
        """
        return self.__append({"render": template})

    def aggregate(self, aggregations):
        """
        Aggregate results of query output
        """
        if not isinstance(aggregations, list):
            aggregations = [aggregations]
        return self.__append({"aggregate": {"aggregations": aggregations}})

    def toJson(self):
        """
        Return the query as a JSON string.
        """
        output = {"query": self.query}
        return json.dumps(output)

    def __iter__(self):
        return self.execute()

    def execute(self):
        """
        Execute the query and return an iterator.
        """
        payload = self.toJson()
        headers = {"Content-Type": "application/json",
                   "Accept": "application/json"}
        request = urllib2.Request(self.url, payload, headers=headers)
        response = do_request(request)
        for result in response:
            try:
                d = json.loads(result)
                if "vertex" in d:
                    yield d["vertex"]
                elif "edge" in d:
                    yield d["edge"]
                elif "aggregations" in d:
                    yield d["aggregations"]["aggregations"]
                elif "selections" in d:
                    d = d["selections"]["selections"]
                    for k in d:
                        if "vertex" in d[k]:
                            d[k] = d[k]["vertex"]
                        elif "edge" in d[k]:
                            d[k] = d[k]["edge"]
                    yield d
                elif "render" in d:
                        yield d["render"]
                elif "count" in d:
                        yield d
                else:
                    yield d
            except ValueError as e:
                print("Can't decode: %s" % (result), file=sys.stderr)
                raise e


class QueryResult(object):

    def __init__(self, data):
        for k in data:
            v = data[k]
            if isinstance(v, dict):
                self.__dict__[k] = self.__class__(v)
            else:
                self.__dict__[k] = v

    def __setattr__(self, k, v):
        if isinstance(v, dict):
            self.__dict__[k] = self.__class__(v)
        else:
            self.__dict__[k] = v

    def __repr__(self):
        attrs = self.as_dict()
        return '<%s %s>' % (self.__class__.__name__, attrs)

    def __str__(self):
        return self.__repr__()

    def as_dict(self):
        attrs = {}
        for a in self.__dict__:
            if not a.startswith('__') and not callable(getattr(self, a)):
                val = getattr(self, a)
                if isinstance(val, dict):
                    for k in val:
                        if isinstance(val[k], QueryResult):
                            attrs[a][k] = val[k].__dict__
                        else:
                            attrs[a] = val
                            break
                elif isinstance(val, QueryResult):
                    attrs[a] = val.__dict__
                else:
                    attrs[a] = val
        return attrs
