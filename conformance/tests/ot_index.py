

def test_label_a_list(O):
    errors = []
    O.addVertexIndex("Person", "name")
    O.addVertex("1", "Person", {"name": "marko", "age": "29"})
    O.addVertex("2", "Person", {"name": "vadas", "age": "27"})
    count = 0
    for i in O.getVertexIndexList():
        count += 1
        if i['field'] != 'name' or i['label'] != 'Person':
            errors.append("Incorrect index field reported")
    if count != 1:
        errors.append("Incorrect number of indices returned")
    return errors


"""
Run label_a_test again, but with different index field to make sure
original index was deleted
"""


def test_label_b_list(O):
    errors = []
    O.addVertexIndex("Person", "age")
    O.addVertex("1", "Person", {"name": "marko", "age": "29"})
    O.addVertex("2", "Person", {"name": "vadas", "age": "27"})
    count = 0
    for i in O.getVertexIndexList():
        count += 1
        if i['field'] != 'age' or i['label'] != 'Person':
            errors.append("Incorrect index field reported")
    if count != 1:
        errors.append("Incorrect number of indices returned")
    return errors


def test_label_index(O):
    errors = []

    O.addVertex("1", "Person", {"name": "marko", "age": "29"})
    O.addVertex("2", "Person", {"name": "vadas", "age": "27"})
    O.addVertex("3", "Software", {"name": "lop", "lang": "java"})
    O.addVertex("4", "Person", {"name": "josh", "age": "32"})
    O.addVertex("5", "Software", {"name": "ripple", "lang": "java"})
    O.addVertex("6", "Person", {"name": "peter", "age": "35"})

    O.addEdge("1", "3", "created", {"weight": 0.4})
    O.addEdge("1", "2", "knows", {"weight": 0.5})
    O.addEdge("1", "4", "knows", {"weight": 1.0})
    O.addEdge("4", "3", "created", {"weight": 0.4})
    O.addEdge("6", "3", "created", {"weight": 0.2})
    O.addEdge("4", "5", "created", {"weight": 1.0})

    res = list(O.query().V().hasLabel("Person").count())[0]
    if res['data'] != 4:
        errors.append(
            "Incorrect vertex index return count: %d != %d" %
            (res['data'], 4))

    res = list(O.query().E().hasLabel("knows").count())[0]
    if res['data'] != 2:
        errors.append(
            "Incorrect edge index return count: %d != %d" %
            (res['data'], 2))

    return errors


def test_count_index(O):
    errors = []

    O.addVertexIndex("Person", "name")

    O.addVertex("1", "Person", {"name": "marko", "age": "29"})
    O.addVertex("2", "Person", {"name": "vadas", "age": "27"})
    O.addVertex("3", "Software", {"name": "lop", "lang": "java"})
    O.addVertex("4", "Person", {"name": "josh", "age": "32"})
    O.addVertex("5", "Software", {"name": "ripple", "lang": "java"})
    O.addVertex("6", "Person", {"name": "peter", "age": "35"})
    O.addVertex("7", "Person", {"name": "marko", "age": "35"})

    O.addEdge("1", "3", "created", {"weight": 0.4})
    O.addEdge("1", "2", "knows", {"weight": 0.5})
    O.addEdge("1", "4", "knows", {"weight": 1.0})
    O.addEdge("4", "3", "created", {"weight": 0.4})
    O.addEdge("6", "3", "created", {"weight": 0.2})
    O.addEdge("4", "5", "created", {"weight": 1.0})

    count = 0
    for row in O.getVertexIndex("Person", "name"):
        count += 1
        if row['term'] == 'marko':
            if row['count'] != 2:
                errors.append(
                    "Incorrect term count: %d != %d" %
                    (row['count'], 2))
        else:
            if row['count'] != 1:
                errors.append(
                    "Incorrect term count: %d != %d" %
                    (row['count'], 1))

    if count != 4:
        errors.append(
            "Incorrect vertex index return count: %d != %d" %
            (count, 4))

    return errors


"""
def test_count_index_query(O):
    errors = []

    O.addVertexIndex("Person", "name")

    O.addVertex("1", "Person", {"name":"marko", "age":"29"})
    O.addVertex("2", "Person", {"name":"vadas", "age":"27"})
    O.addVertex("3", "Software", {"name":"lop", "lang":"java"})
    O.addVertex("4", "Person", {"name":"josh", "age":"32"})
    O.addVertex("5", "Software", {"name":"ripple", "lang":"java"})
    O.addVertex("6", "Person", {"name":"peter", "age":"35"})
    O.addVertex("7", "Person", {"name":"marko", "age":"35"})

    O.addEdge("1", "3", "created", {"weight":0.4})
    O.addEdge("1", "2", "knows", {"weight":0.5})
    O.addEdge("1", "4", "knows", {"weight":1.0})
    O.addEdge("4", "3", "created", {"weight":0.4})
    O.addEdge("6", "3", "created", {"weight":0.2})
    O.addEdge("4", "5", "created", {"weight":1.0})

    count = 0
    for row in O.index().query("Person", "name", "marko").outgoing():
        count += 1
    if count != 2:
        errors.append("Wrong index query return count: %s != %s" % (count, 2))

    return errors
"""
