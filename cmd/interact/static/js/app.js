document.addEventListener('DOMContentLoaded', (event) => {
  console.log('DOM fully loaded and parsed');
  searchHandler();
});

function searchHandler() {
  let s = document.getElementById("search");
  s.addEventListener("input", () => createSelection(autoComplete(s.value)));
}
