
def test_fields(O):
    errors = []

    O.addVertex("vertex1", "person", {"name": "han", "age": 35, "occupation": "smuggler", "starships": ["millennium falcon"]})

    expected = {
        u"gid": u"vertex1",
        u"label": u"",
        u"data": {u"name": u"han"}
    }
    resp = O.query().V().fields(["_gid", "name"]).execute()
    if resp[0] != expected:
        errors.append("vertex contains incorrect fields: %s" % resp)

    expected = {
        u"gid": u"vertex1",
        u"label": u"",
        u"data": {u"non-existent": None}
    }
    resp = O.query().V().fields(["_gid", "non-existent"]).execute()
    if resp[0] != expected:
        errors.append("vertex contains incorrect fields: %s" % resp)

    return errors
