function autoComplete(str) {
  if (!str) {
    console.log('autocomplete called with empty string');
    return []
  }
  return [str,'a','b','c']
}
