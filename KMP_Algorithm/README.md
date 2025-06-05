# KMP (Knuth-Morris-Pratt) Algorithm

The KMP algorithm is an efficient string matching algorithm that finds occurrences of a "pattern" string within a "text" string. It uses the pattern's structure to minimize comparisons by avoiding unnecessary backtracking.

## Overview

The KMP algorithm improves upon naive string matching by using a prefix table (also called LPS - Longest Proper Prefix which is also Suffix) to skip redundant comparisons when a mismatch occurs.

### Time Complexity
- Preprocessing (building LPS array): O(m)
- Pattern matching: O(n)
- Total: O(n + m)
where n is the length of text and m is the length of pattern

### Space Complexity
- O(m) for the LPS array
- O(1) additional space for searching

## Implementation Methods

This implementation includes several approaches:

1. **Prefix Table (LPS Array)**
   - Core component of KMP
   - Tracks longest proper prefix that is also a suffix
   - Used to skip unnecessary comparisons

2. **Basic KMP Search**
   - Finds all non-overlapping occurrences
   - Returns array of match indices
   - Optimal for standard pattern matching

3. **First Occurrence Search**
   - Stops at first match
   - Returns single index
   - Efficient for existence checking

4. **Overlapping Search**
   - Finds overlapping matches
   - Useful for special cases
   - Example: finding "AA" in "AAA"

## Usage Examples

```go
// Building the prefix table
pattern := "AABA"
lps := kmp.ComputeLPSArray(pattern)
// lps = [0, 1, 0, 1]

// Finding all occurrences
text := "AABAACAADAABAAABAA"
matches := kmp.KMPSearch(text, pattern)
// matches = [0, 9, 13]

// Finding first occurrence
firstMatch := kmp.KMPSearchFirst(text, pattern)
// firstMatch = 0

// Counting occurrences
count := kmp.KMPCount(text, pattern)
// count = 3
```

## Understanding the Prefix Table (LPS Array)

The LPS array is crucial for KMP's efficiency. For each position i, LPS[i] contains the length of the longest proper prefix that is also a suffix of the pattern[0...i].

Example for pattern "AABAACAABAA":
```
Pattern: A  A  B  A  A  C  A  A  B  A  A
LPS:     0  1  0  1  2  0  1  2  3  4  5
```

This means:
- LPS[4] = 2 because "AA" is the longest prefix also suffix of "AABAA"
- LPS[10] = 5 because "AABAA" is the longest prefix also suffix of "AABAACAABAA"

## Applications

1. **String Matching**
   - Finding substring occurrences
   - DNA sequence matching
   - Plagiarism detection

2. **Pattern Analysis**
   - Text processing
   - Data compression
   - Network protocols

3. **Biological Computing**
   - Gene sequence matching
   - Protein pattern matching
   - Molecular biology

## Edge Cases Handled

- Empty text or pattern
- Pattern longer than text
- Single character text/pattern
- Overlapping patterns
- Repeated characters

## Performance Considerations

1. **Preprocessing**
   - One-time LPS array computation
   - Reusable for same pattern

2. **Memory Usage**
   - Only stores pattern-length array
   - Constant extra space for search

3. **Comparison Count**
   - Never compares same text character twice
   - Optimal for worst-case scenarios

## Common Use Cases

1. **Text Editors**
   - Find/Replace functionality
   - Pattern highlighting

2. **Bioinformatics**
   - DNA sequence alignment
   - Protein matching

3. **Data Mining**
   - Pattern frequency analysis
   - Text classification

4. **Network Security**
   - Intrusion detection
   - Pattern-based filtering 