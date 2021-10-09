//------------------------------------------------------------------------------
// Problem 1.1

const allUnique = (str) => {
  const set = new Set();
  for (const ch of str) {
    if (set.has(ch)) return false;
    set.add(ch);
  }
  return true;
};

const allUniqueNoSet = (str) => {
  const array = [];
  for (let i = 0; i < str.length; i++) {
    if (array[str.charCodeAt(i)]) return false;
    array[str.charCodeAt(i)] = true;
  }
  return true;
};
