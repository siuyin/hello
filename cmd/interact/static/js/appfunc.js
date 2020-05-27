function createSelection(list) {
  // creates a HTML select with options given in list. list is an array of strings.
  let ss = document.getElementById("search-select");
  if (list.length == 0) {
    ss.style.display = "none";
    return;
  }
  ss.style.display = "block";
  console.log('calling auto complete with: '+list+'.');
}
function autoComplete(str) {
  if (!str) {
    console.log('autocomplete called with empty string');
    return []
  }
  return [str,'a','b','c']
}
