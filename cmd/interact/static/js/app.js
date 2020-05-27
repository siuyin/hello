document.addEventListener('DOMContentLoaded', (event) => {
  console.log('DOM fully loaded and parsed');
  searchHandler();
});

function searchHandler() {
  let s = document.getElementById("search");
  s.addEventListener("input", () => createSelection(autoComplete(s.value)));
}
function createSelection(list) {
  // creates a HTML select with options given in list. list is an array of strings.
  console.log('calling auto complete with: '+list+'.');
}
function autoComplete(str) {
  console.log('autocomplete called with: '+str);
  return [str]
}
