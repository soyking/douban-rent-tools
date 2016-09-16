// 对关键字进行扩充
// 如果关键字满足一个 Condition 则进行 Expand 扩充
package expand

// 条件判断
type Condition func(string) bool

// 扩充，返回扩充结果添加到原来关键字
type Expand func(string) []string

type Expander struct {
	conditions []Condition
	expands    []Expand
}

func NewExpander() *Expander {
	return &Expander{}
}

func (e *Expander) AddExpand(condition Condition, expand Expand) *Expander {
	if condition != nil && expand != nil {
		e.conditions = append(e.conditions, condition)
		e.expands = append(e.expands, expand)
	}

	return e
}

func (e *Expander) Expand(keywords []string) []string {
	expandedKeywords := []string{}
	for _, keyword := range keywords {
		meetCond := false
		for i, cond := range e.conditions {
			if cond(keyword) {
				// keywords[i] 也想加入的话要在 expand 中加入
				meetCond = true
				expandedKeywords = append(expandedKeywords, e.expands[i](keyword)...)
			}
		}
		if !meetCond {
			expandedKeywords = append(expandedKeywords, keyword)
		}
	}
	return expandedKeywords
}
