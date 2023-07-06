import weaviate
import json

client = weaviate.Client(
    url="https://some-endpoint.weaviate.network",  # Replace with your endpoint
    auth_client_secret=weaviate.AuthApiKey(
        api_key="YOUR-WEAVIATE-API-KEY"
    ),  # Replace w/ your Weaviate instance API key
)

# ===== Create schema =====
# Class definition object. Weaviate's autoschema feature will infer properties when importing.
class_obj = {
    "class": "Question",
    "vectorizer": "text2vec-openai",
}

# Add the class to the schema
client.schema.create_class(class_obj)

# ===== Import data =====
import requests

url = "https://raw.githubusercontent.com/weaviate-tutorials/quickstart/main/data/jeopardy_tiny+vectors.json"
resp = requests.get(url)
data = json.loads(resp.text)

# Configure a batch process
with client.batch as batch:
    batch.batch_size = 100
    # Batch import all Questions
    for i, d in enumerate(data):
        print(f"importing question: {i+1}")

        properties = {
            "answer": d["Answer"],
            "question": d["Question"],
            "category": d["Category"],
        }

        client.batch.add_data_object(properties, "Question", vector=d["Vector"])
