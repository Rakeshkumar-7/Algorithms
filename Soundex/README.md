# Soundex algorithm - Phonetic fuzzy search

Matches two strings based on their sound similarity instead of plain old [[Levenshtein Distance]]. Check out more [here](https://stackoverflow.com/questions/32337135/fuzzy-search-algorithm-approximate-string-matching-algorithm) (Read Jim Mischel's response)

Pseudo-code ([wikipedia](https://en.wikipedia.org/wiki/Soundex), [stanford](https://web.stanford.edu/class/archive/cs/cs106b/cs106b.1216/assignments/1-cpp/soundex)):
1. Replace (a, e, i, o, u, y, h, w) with 0
2. Replace consonants with these numeric value:
	- b, f, p, v → 1
	- c, g, j, k, q, s, x, z → 2
	- d, t → 3
	- l → 4
	- m, n → 5
	- r → 6
3.  Coalesce adjacent duplicate digits
4. Remove all zeros
5. Replace first number with the upper-case equivalent of the first letter of the input
6. Include the numbers from index=1 till the length of the output becomes 4
7. If there are not enough digits, pad it with 0 till the length is 4


**Note:** Works with ASCII and English.

**Sample Output:**
hello: H400
hallow: H400
test: T230
smith: S530
smythe: S530
smote: S530