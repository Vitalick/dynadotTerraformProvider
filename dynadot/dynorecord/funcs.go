package dynorecord

func expandStringList(configured []interface{}) []string {
	var result []string
	for _, raw := range configured {
		result = append(result, raw.(string))
	}
	return result
}
