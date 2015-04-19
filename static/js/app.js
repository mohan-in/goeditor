
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

document.querySelector('#btn-dir').addEventListener('click', function(e) {
	e.preventDefault();
	document.querySelector('#wrapper').classList.toggle('toggled');
});

function registerFileButtonEvents(btn) {
	btn.addEventListener('click', function() {
		var childs = document.querySelector('#btn-files').querySelectorAll("*")
		for (var i = childs.length - 1; i >= 0; i--) {
			childs[i].classList.remove('btn-primary');
			childs[i].classList.add('btn-default');
		};

		btn.classList.add('btn-primary');
		btn.classList.remove('btn-default');
	});
}

registerFileButtonEvents(document.querySelector('#a-go'));
registerFileButtonEvents(document.querySelector('#b-go'));
registerFileButtonEvents(document.querySelector('#c-go'));


function getTree() {
	var tree = [
  {
    text: "Parent 1",
    nodes: [
      {
        text: "Child 1",
        nodes: [
          {
            text: "Grandchild 1"
          },
          {
            text: "Grandchild 2"
          }
        ]
      },
      {
        text: "Child 2"
      }
    ]
  },
  {
    text: "Parent 2"
  },
  {
    text: "Parent 3"
  },
  {
    text: "Parent 4"
  },
  {
    text: "Parent 5"
  }
];
	return tree;
}

$('#dir-tree').treeview({data: getTree(), selectable: false});
