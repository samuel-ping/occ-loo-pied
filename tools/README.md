# Tools

## generateSampleMetrics.py

### Usage

```
python generateSampleMetrics.py
```

This will write a JSON file `sampleMetrics.json` into the same directory. You can load this file's data into your MongoDB collection by running the following command:

```
cat sampleMetrics.json | mongoimport --collection='occloopied.metrics'

# For MongoDB instances running in Docker
docker exec -i occ-loo-pied-db-1 sh -c 'mongoimport -c metrics -d occloopied --jsonArray --uri="mongodb://root:password@localhost:27017/occloopied?authSource=admin"' < sampleMetrics.json
```