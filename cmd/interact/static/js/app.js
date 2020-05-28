document.addEventListener('DOMContentLoaded', (event) => {
  console.log('DOM fully loaded and parsed');
  searchHandler();
});

function searchHandler() {
  let s = document.getElementById("search");
  s.addEventListener("input", () => createSelection(autoCompleteAPI(s.value)));
}
async function createSelection(listPromise) {
  // creates a HTML select with options given in list. list is an array of strings.
  let ss = document.getElementById("search-select");
  const list = await listPromise;

  if (!list || list.length == 0) {
    ss.style.display = "none";
    return;
  }
  console.log(`creating ${list}`);
  csCreate(list);
  ss.style.display = "block";
}
function csCreate(list) {
  let ss = document.getElementById("search-select");
  ss.innerHTML = '';

  for (x of list) {
    let o = document.createElement("option");
    o.setAttribute("value",x);
    o.innerHTML = x;
    ss.appendChild(o);
  }
}

async function autoCompleteAPI(str) {
  let f = await fetch("/bar")
  .then(response => response.text());

  if (!str) {
    return [];
  }
  return [f,str];
}
