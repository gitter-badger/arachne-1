function query() {
	function process(l) {
		if (!l) {
			l = []
		} else if (_.isString(l)) {
			l = [l]
		} else if (!_.isArray(l)) {
			throw "not something we know how to make labels out of"
		}
		return l
	}

	return {
		query: [],
		V: function(id) {
			this.query.push({'v': process(id)})
			return this
		},
		E: function(id) {
			this.query.push({'e': process(id)})
			return this
		},
		out: function(label) {
			this.query.push({'out': process(label)})
			return this
		},
		in_: function(label) {
			this.query.push({'in': process(label)})
			return this
		},
		both: function(label) {
			this.query.push({'both': process(label)})
			return this
		},
		outEdge: function(label) {
			this.query.push({'out_edge': process(label)})
			return this
		},
		inEdge: function(label) {
			this.query.push({'in_edge': process(label)})
			return this
		},
		bothEdge: function(label) {
			this.query.push({'both_edge': process(label)})
			return this
		},
		mark: function(name) {
			this.query.push({'mark': name})
			return this
		},
		select: function(marks) {
			this.query.push({'select': {'marks': process(marks)}})
			return this
		},
		limit: function(n) {
			this.query.push({'limit': n})
			return this
		},
		offset: function(n) {
			this.query.push({'offset': n})
			return this
		},
		count: function() {
			this.query.push({'count': ''})
			return this
		},
		distinct: function(val) {
			this.query.push({'distinct': process(val)})
			return this
		},
		render: function(r) {
			this.query.push({'render': r})
			return this
		},
		where: function(expression) {
			this.query.push({'where': expression})
			return this
		},
		aggregate: function() {
			this.query.push({'aggregate': {'aggregations': Array.prototype.slice.call(arguments)}})
			return this
		}
	}
}

// Where operators
function and_() {
	return {'and': {'expressions': Array.prototype.slice.call(arguments)}}
}

function or_() {
	return {'or': {'expressions': Array.prototype.slice.call(arguments)}}
}

function not_(expression) {
	return {'not': expression}
}

function eq(key, value) {
	return {'condition': {'key': key, 'value': value, 'condition': 'EQ'}}
}

function neq(key, value) {
	return {'condition': {'key': key, 'value': value, 'condition': 'NEQ'}}
}

function gt(key, value) {
	return {'condition': {'key': key, 'value': value, 'condition': 'GT'}}
}

function gte(key, value) {
	return {'condition': {'key': key, 'value': value, 'condition': 'GTE'}}
}

function lt(key, value) {
	return {'condition': {'key': key, 'value': value, 'condition': 'LT'}}
}

function lte(key, value) {
	return {'condition': {'key': key, 'value': value, 'condition': 'LTE'}}
}

function in_(key, values) {
	if (!values) {
		values = []
	} else if (!_.isObject(l) && !_.isArray(l)) {
		values = [values]
	}
	return {'condition': {'key': key, 'value': values, 'condition': 'IN'}}
}

function contains(key, value) {
	return {'condition': {'key': key, 'value': value, 'condition': 'CONTAINS'}}
}

// Aggregation builders
function term(name, label, field, size) {
	agg = {
		"name": name,
		"term": {"label": label, "field": field}
	}
	if (size) {
		if (!_.isNumber(percents)) {
			throw "size expected to be a number"
		}
		agg["term"]["size"] = size
	}
	return agg
}

function percentile(name, label, field, percents) {
	if (!percents) {
		percents = [1, 5, 25, 50, 75, 95, 99]
	} else if (_.isNumber(percents)) {
			percents = [percents]
	} else if (!_.isArray(percents)) {
		throw "percents expected to be an array of numbers"
	}

	return {
		"name": name,
		"percentile": {
			"label": label, "field": field, "percents": percents
		}
	}
}

function histogram(name, label, field, interval) {
	if (interval) {
		if (!_.isNumber(interval)) {
			throw "size expected to be a number"
		}
	}
	return {
		"name": name,
		"histogram": {
			"label": label, "field": field, "interval": interval
		}
	}
}

// base query object
O = {
	query : query
}
