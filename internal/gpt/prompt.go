package gpt

type TitleBody struct {
	Title string `json:"title"`
}
type ContextBody struct {
	Context string `json:"context"`
}

var GET_TITLE = "```%s``` 해당 HTML에서 포스팅 제목을 json 형식으로만 간단하게 출력해줘"

var POST_SUMMARY = "```%s``` 해당 HTML에서 포스팅을 한글로 간단하게 요약한 결과를 {'context':string} 과 같은 json 형식으로만 간단하게 출력해줘"
