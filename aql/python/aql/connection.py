from __future__ import absolute_import, print_function, unicode_literals

import json
import requests


from aql.util import process_url
from aql.graph import Graph


class Connection:
    def __init__(self, url):
        url = process_url(url)
        self.base_url = url
        self.url = url + "/v1/graph"

    def listGraphs(self):
        """
        List graphs.
        """
        response = requests.get(self.url, stream=True)
        response.raise_for_status()
        output = []
        for line in response.iter_lines():
            output.append(json.loads(line)['graph'])
        return output

    def addGraph(self, name):
        """
        Create a new graph.
        """
        response = requests.post(self.url + "/" + name, {})
        response.raise_for_status()
        return response.json()

    def deleteGraph(self, name):
        """
        Delete graph.
        """
        response = requests.delete(self.url + "/" + name)
        response.raise_for_status()
        return response.json()

    def graph(self, name):
        """
        Get a graph handle.
        """
        return Graph(self.base_url, name)
