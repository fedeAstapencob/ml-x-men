package testData

var PersonDefinition = `
"Person": {
  "type": "object",
  "properties": {
	"dna": {
      "type": "string"
    },
    "isMutant": {
      "type": "boolean"
    }
  }
}`

var PersonRespDefinition = `{
	` + PersonDefinition + `,
      "type": "object",
      "properties": {
        "person": {
          "$ref": "#/Person"
        }
      }
}`

var StatsRespDefinition = `{
	"type": "object",
	"properties": {
		"count_mutant_dna": {
			"type": "number"
		},
		"count_human_dna": {
			"type": "number"
		},
		"ratio": {
			"type": "number"
		}
	} 
}`
