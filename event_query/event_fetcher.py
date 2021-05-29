import sys
import requests
import json
from flask import Flask
import os

app = Flask(__name__)

@app.route('/events')
def event_querier():
    apiserver = os.environ['KUBE_API_SERVER']
    token = os.environ['KUBE_API_SERVER_TOKEN']
    headers = {"Authorization":f"Bearer {token}"}
    result = requests.get(f"{apiserver}/api/v1/events",headers=headers,verify=False)
    event_list = []
    for item in  result.json()['items']:

        event_list.append({
             'eventname' : item['metadata']['name'],
             'event_creationTimestamp': item['metadata']['creationTimestamp'],
             'kind' : item['involvedObject']['kind'],
             'namespace':item['involvedObject']['namespace'],
             'name':item['involvedObject']['name'],
             'message':item['message'],
             'reason':item['reason'],
             'firstTimestamp':item['firstTimestamp'],
             'lastTimestamp':item['lastTimestamp'],
             'typename':item['type']
        })

    return json.dumps(event_list)


if __name__ == '__main__':
    args = sys.argv
    app.run(host='0.0.0.0',debug=True, port=5000)




        