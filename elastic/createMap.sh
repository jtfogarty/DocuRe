curl -XPUT '192.168.1.149:9200/shakespeare?pretty' -H 'Content-Type: application/json' -d'
{
 "mappings": {
  "doc": {
   "properties": {
    "speaker": {"type": "keyword"},
    "play_name": {"type": "keyword"},
    "line_id": {"type": "integer"},
    "speech_number": {"type": "integer"}
   }
  }
 }
}
'


curl -XPUT '192.168.1.149:9200/american_standard_version?pretty' -H 'Content-Type: application/json' -d'
{
 "mappings": {
  "doc": {
   "properties": {
    "book": {"type": "keyword"},
    "chapter_no": {"type": "integer"},
    "verse_no": {"type": "integer"},    
    "line_id": {"type": "integer"}
   }
  }
 }
}
'

curl -XPUT '192.168.1.149:9200/king_james_bible?pretty' -H 'Content-Type: application/json' -d'
{
 "mappings": {
  "doc": {
   "properties": {
    "book": {"type": "keyword"},
    "chapter_no": {"type": "integer"},
    "verse_no": {"type": "integer"},    
    "line_id": {"type": "integer"}
   }
  }
 }
}
'

curl -XPUT '192.168.1.149:9200/american_king_james?pretty' -H 'Content-Type: application/json' -d'
{
 "mappings": {
  "doc": {
   "properties": {
    "book": {"type": "keyword"},
    "chapter_no": {"type": "integer"},
    "verse_no": {"type": "integer"},    
    "line_id": {"type": "integer"}
   }
  }
 }
}
'

curl -XPUT '192.168.1.149:9200/youngs_literal_translation?pretty' -H 'Content-Type: application/json' -d'
{
 "mappings": {
  "doc": {
   "properties": {
    "book": {"type": "keyword"},
    "chapter_no": {"type": "integer"},
    "verse_no": {"type": "integer"},    
    "line_id": {"type": "integer"}
   }
  }
 }
}
'


curl -XPUT '192.168.1.149:9200/jerusalem_bible?pretty' -H 'Content-Type: application/json' -d'
{
 "mappings": {
  "doc": {
   "properties": {
    "book": {"type": "keyword"},
    "chapter_no": {"type": "integer"},
    "verse_no": {"type": "integer"},    
    "line_id": {"type": "integer"}
   }
  }
 }
}
'
