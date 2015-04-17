
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

document.querySelector('#btn-dir').addEventListener('click', function() {
	editor.focus();
});

document.querySelector('#btn-save').addEventListener('click', function() {
	editor.focus();
});

document.querySelector('#btn-fmt').addEventListener('click', function() {
	editor.focus();
});

