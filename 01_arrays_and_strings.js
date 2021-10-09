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
