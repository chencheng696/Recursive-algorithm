func CreateTreeData_(parentNo string, src []obj.GroupInfoma)[]interface{} {
	arr := make([]interface{}, 0)
	
  for _, item := range src {
		if item.Id == parentNo{
			var numberString string = fmt.Sprintf("(%s)", strconv.Itoa(item.ChildNum))
			var titleString string = item.Name + numberString

			tree := map[string]interface{}{
				`key`       :            item.Id,
				`title`     :            titleString,
				`parent`    :            item.Parent,
				`isLeaf`    :            item.IsLeaf,
				`children`  :            CreateTreeData(item.Id, src),
			}

			arr = append(arr, tree)
		}
	
	return arr
}
