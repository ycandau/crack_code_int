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

//------------------------------------------------------------------------------
// Problem 1.2

const checkPermutation = (str1, str2) => {
  const counts = new Map();
  for (let i = 0; i < str1.length; i++) {
    const ch = str1[i];
    const count = counts.get(ch) || 1;
    counts.set(ch, count);
  }
  for (let i = 0; i < str2.length; i++) {
    const ch = str2[i];
    if (!counts.has(ch)) return false;
    counts.delete(ch);
  }
  if (counts.size !== 0) return false;
  return true;
};
