
var editor = CodeMirror.fromTextArea(document.querySelector('#editor'), {
	theme: "elegant",
	lineNumbers: true,
	matchBrackets: true,
	indentUnit: 4,
	tabSize: 4,
	indentWithTabs: true,
	mode: "go",
	autofocus: true
});

