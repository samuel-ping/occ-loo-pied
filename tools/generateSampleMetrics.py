import json
import random
from datetime import datetime, timedelta

'''
Generates sample metrics documents to a JSON file that can be inserted into MongoDB.
Sample document:
{
  endTime: ISODate('2025-02-01T00:50:04.546Z'),
  duration: Long('568591891'),
  startTime: ISODate('2025-02-01T00:50:03.977Z')
}
'''
def generate_documents(days: int):
    documents = []
    now = datetime.now()
    
    for day_offset in range(days):
        date = now - timedelta(days=day_offset)
        num_entries = random.randint(0, 50)
        
        times_used = set()
        
        for _ in range(num_entries):
            while True:
                start_time = datetime(
                    date.year, date.month, date.day,
                    random.randint(0, 23), random.randint(0, 59), random.randint(0, 59)
                )
                if start_time not in times_used:
                    break
            
            duration = random.randint(1, 3600) * 1000  # Duration in milliseconds (up to 1 hour)
            end_time = start_time + timedelta(milliseconds=duration)
            times_used.add(start_time)
            times_used.add(end_time)

            documents.append({
                "startTime": {"$date": start_time.isoformat() + "Z"},
                "endTime": {"$date": end_time.isoformat() + "Z"},
                "duration": {"$numberLong": str(duration)}
            })
    
    return documents

# Generate data for the past 30 days
data = generate_documents(30)

# Save to a JSON file
with open("sampleMetrics.json", "w") as f:
    json.dump(data, f, indent=2)

print(f"Generated {len(data)} documents.")
