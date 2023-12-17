let content = "";

const dropzoneArea = document.getElementById("drop-area");
const hiddenDropzone = document.getElementById("file-input");
const errorDropzone = document.getElementById("dropzone-error");

const inputArea = document.getElementById("content");
const contentLength = document.getElementById("content-length");

dropzoneArea.addEventListener("click", () => hiddenDropzone.click());
hiddenDropzone.addEventListener("change", (e) => handleDrop(e, true));
inputArea.addEventListener("input", () => {
    content = inputArea.value;
    contentLength.innerText = content.length;
});

const all_events = ["dragenter", "dragover", "dragleave", "drop"];
const enter_events = ["dragenter", "dragover"];
const leave_events = ["dragleave", "drop"];

for (let event of all_events) {
    dropzoneArea.addEventListener(event, (e) => {
        e.preventDefault();
        e.stopPropagation();
    });
}

for (let event of enter_events) {
    dropzoneArea.addEventListener(event, highlight);
}

for (let event of leave_events) {
    dropzoneArea.addEventListener(event, unhighlight);
}

function highlight() {
    dropzoneArea.classList.remove("border-slate-500");
    dropzoneArea.classList.add("border-black");
}

function unhighlight() {
    dropzoneArea.classList.add("border-slate-500");
    dropzoneArea.classList.remove("border-black");
}

dropzoneArea.addEventListener("drop", (e) => handleDrop(e, false));

function handleDrop(e, isClick) {
    errorDropzone.innerText = "";
    let file = isClick ? e.target.files[0] : e.dataTransfer.files[0];
    if (file.type !== "text/plain" && file.type !== "text/markdown") return errorDropzone.innerText = "Your file is not .txt or .md";
    parseFile(file);
}

function parseFile(file) {
    const reader = new FileReader();

    reader.addEventListener("load", (event) => {
        const result = event.target.result;
        if (result !== null && result !== "") {
            content = result;
            updateInput();
        }
    });

    reader.readAsText(file);
}

function updateInput() {
    inputArea.value = content;
    contentLength.innerText = content.length;
}
