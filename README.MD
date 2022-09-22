<!-- ABOUT THE PROJECT -->

# About the project

Leetcode & DailyCodingProblem practice using GoLang

![Build Status](https://github.com/serhii-soboliev/golc/actions/workflows/go.yml/badge.svg)
![Code Quality Score](https://api.codiga.io/project/34041/score/svg)
![Code Grade](https://api.codiga.io/project/34041/status/svg)

### Table of Contents

**[Leetcode](#leetcode-problems)** \
**[Daily Coding Problem](#daily-coding-problem)**

## Leetcode Problems

|#  | Problem          | Link To Solution                              | Method               |
|:-:| :---             | :-                                            | :-                   |
|[1](https://leetcode.com/problems/two-sum/)| Two Sum | [link](pkg/leetcode/hashmap/1.two_sum.go)| HashMap |
|[2](https://leetcode.com/problems/add-two-numbers/)| Add Two Numbers | [link](pkg/leetcode/linkedlist/2.add_two_numbers.go)| Linked List|
|[3](https://leetcode.com/problems/longest-substring-without-repeating-characters/)| Longest Substring Without Repeating Characters| [link](pkg/leetcode/slidingwindow/3.longest_substring_wo_rep_char.go)| Sliding Window|
|[4](https://leetcode.com/problems/median-of-two-sorted-arrays/)| 4. Median of Two Sorted Arrays| [link](pkg/leetcode/binarysearch/4.median_of_two_sorted_arrays.go)| Binary Search |
|[5](https://leetcode.com/problems/longest-palindromic-substring/)| Longest Palindromic Substring | [link](pkg/leetcode/string/5.longest_palindromic_substring.go)| String|
|[6](https://leetcode.com/problems/zigzag-conversion/)| Zigzag Conversion | [link](pkg/leetcode/string/6.zigzag_conversion.go)| String|
|30| Substring with Concatenation of All Words | [link](pkg/leetcode/string/30.substring_with_concatenation_of_all_words.go)| String|
|33| Sudoku Solver | [link](pkg/leetcode/backtracking/33.sudoku_solver.go)| Backtracking|
|[41](https://leetcode.com/problems/first-missing-positive/)| First Missing Positive | [link](pkg/dailycodingproblems/solutions/stripe/problem_4.go)| Array|
|[42](https://leetcode.com/problems/trapping-rain-water/)| Trapping Rain Water | [link](pkg/leetcode/dynamicprogramming/42.trapping_rain_water.go)| DP, Brute Force, Stack|
|48| Rotate Image | [link](pkg/leetcode/matrix/48.rotate_image.go)| Matrix|
|66| Plus One | [link](pkg/leetcode/math/66.plus_one.go)| Math |
|86| Partition list | [link](pkg/leetcode/linkedlist/86.partition_list.go)| Linked List|
|122| Word Ladder II | [link](pkg/leetcode/backtracking/126.word_ladder_2.go)| Backtracking|
|140| Word Break II | [link](pkg/leetcode/backtracking/140.word_break_2.go)| Backtracking|
|[188](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/)| Best Time to Buy and Sell Stock IV | [link](pkg/leetcode/dynamicprogramming/188.best_time_buy_sell_stock4.go)| DP |
|199| Binary Tree Right Side View | [link](pkg/leetcode/tree/199.binary_tree_right_side_view.go)| Binary Tree|
|208| Trie(Prefix tree) | [link](pkg/leetcode/datastructures/trie/208.trie.go)| Trie|
|212| Word Search 2 | [link](pkg/leetcode/backtracking/212.word_search_2.go)| Backtracking |
|234| Palindrome Linked List | [link](pkg/leetcode/linkedlist/234.palindrome_linked_list.go)| Linked List |
|[238](https://leetcode.com/problems/product-of-array-except-self/)| Product of Array Except Self | [link](pkg/dailycodingproblems/uber/problem_2.go)| Array |
|240| Search 2D Matrix 2 | [link](pkg/leetcode/binarysearch/240.search_2d_matrix_ll.go)| Binary Search |
|242| Valid Anagram | [link](pkg/leetcode/string/242.valid_anagram.go)| Binary Search |
|282| Expression And Operators| [link](pkg/leetcode/backtracking/282.expression_add_operators.go)| Backtracking |
|[297](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)| Serialize and Deserialize Binary Tree| [link](pkg/dailycodingproblems/solutions/google/problem_3.go)| Data structure |
|301| Remove Invalid Parentheses| [link](pkg/leetcode/backtracking/301.remove_invalid_parentheses.go)| Backtracking |
|363| Max Sum of Rectangle No Larger Than K| [link](pkg/leetcode/matrix/363.max_sum_rectangle_no_larger_k.go)| Prefix Sum |
|377| Combination Sum IV | [link](pkg/leetcode/dynamicprogramming/377.combination_sum_4.go)| DP |
|429| N-ary Tree Level Order Traversal| [link](pkg/leetcode/tree/429.nary_tree_level_order_traversal.go)| Tree |
|458| Poor Pigs | [link](pkg/leetcode/math/458.poor_pigs.go)| Math |
|491| Increasing Subsequences| [link](pkg/leetcode/backtracking/491.increasing_subsequences.go)| String|
|494| Target Sum | [link](pkg/leetcode/backtracking/494.target_sum.go)| Backtracking|
|526| Beautiful Arrangement | [link](pkg/leetcode/backtracking/526.beautiful_arrangement.go)| Backtracking|
|[526](https://leetcode.com/problems/reverse-words-in-a-string-iii/)| Reverse Words in a String III | [link](pkg/leetcode/twopointers/557.reverse_words_str3.go)| Two pointers|
|[606](https://leetcode.com/problems/construct-string-from-binary-tree/)| Construct String from Binary Tree  | [link](pkg/leetcode/tree/606.construct_str_from_binary_tree.go)| Tree|
|[609](https://leetcode.com/problems/find-duplicate-file-in-system/)| Find Duplicate File in System  | [link](pkg/leetcode/hashmap/609.duplicate_fill.go)| HashMap|
|638| Shopping Offers  | [link](pkg/leetcode/backtracking/638.shopping_offers.go)| Backtracking|
|659| Split Array into Consecutive Subsequences  | [link](pkg/leetcode/greedy/659.split_array_into_consecutive_subsequences.go)| Greedy|
|679| 24 Game  | [link](pkg/leetcode/backtracking/679.24_game.go)| Backtracking|
|691| Stickers to Spell Word | [link](pkg/leetcode/backtracking/691.stickers_to_spell_word.go)| Backtracking|
|698| Partition to K Equal Sum Subsets  | [link](pkg/leetcode/backtracking/698.partition_k_equal_sum_subsets.go)| Backtracking |
|[718](https://leetcode.com/problems/maximum-length-of-repeated-subarray/)| Maximum Length of Repeated Subarray  | [link](pkg/leetcode/dynamicprogramming/719.max_len_rep_subarray.go)| DP |
|784| Letter Case Permutation  | [link](pkg/leetcode/backtracking/784.letter_case_permutation.go)| Backtracking |
|792| Number of Matching Subsequences  | [link](pkg/leetcode/string/792.number_of_matching_subsequences.go)| String|
|804| Unique Morse Code Words | [link](pkg/leetcode/string/804.unique_morse_code_words.go)| HashTable|
|[814](https://leetcode.com/problems/binary-tree-pruning/)| Binary Tree Pruning | [link](pkg/leetcode/tree/814.binary_tree_pruning.go)| Tree|
|816| Ambiguous Coordinates | [link](pkg/leetcode/backtracking/816.ambiguous_coordinates.go)| Backtracking|
|823| Binary Trees With Factors | [link](pkg/leetcode/dynamicprogramming/823.binary_trees_with_factors.go)| DP|
|842| Split Array into Fibonacci Sequence | [link](pkg/leetcode/backtracking/842.split_array_into_fibonacci_sequence.go)| Backtracking|
|858| Mirror Reflection| [link](pkg/leetcode/math/858.mirror_reflection.go)| Math|
|871| Minimum Number of Refueling Stops| [link](pkg/leetcode/dynamicprogramming/871.minimum_number_of_refst.go)| DP, Priority Queue|
|890| Find and Replace Pattern| [link](pkg/leetcode/string/890.find_and_replace_pattern.go)| String|
|936| Stamping The Sequence| [link](pkg/leetcode/greedy/936.stamping_the_sequence.go)| Greedy|
|[948](https://leetcode.com/problems/bag-of-tokens/)| Bag of Tokens| [link](pkg/leetcode/greedy/936.stamping_the_sequence.go)| Greedy|
|967| Numbers With Same Consecutive Differences| [link](pkg/leetcode/backtracking/967.numbers_with_same_consecutive_diff.go)| Backtracking|
|980| Unique Paths III| [link](pkg/leetcode/backtracking/980.unique_paths_III.go)| Backtracking|
|[985](https://leetcode.com/problems/sum-of-even-numbers-after-queries/)| Sum of Even Numbers After Queries | [link](pkg/leetcode/simulation/985.sum_even_numbers_after_q.go)| Simulation|
|996| Number of Squareful Arrays| [link](pkg/leetcode/backtracking/996.number_of_squareful_arrays.go)| Backtracking|
|1079| Letter Tile Possibilities| [link](pkg/leetcode/backtracking/1079.letter_tile_possibilities.go)| Backtracking |
|1096| Brace Expansion II| [link](pkg/leetcode/bfs/1096.brace_expansion_II.go)| BFS |
|1307| Verbal Arithmetic Puzzle| [link](pkg/leetcode/backtracking/1307.verbal_arithmetic_puzzle.go)| Backtracking |
|1338| Reduce Array Size to The Half| [link](pkg/leetcode/hashmap/1338.reduce_array_size_to_the_half.go)| HashMap |
|1220| Count Vowels Permutation| [link](pkg/leetcode/dynamicprogramming/1220.count_vowels_permutation.go)| DP|
|1255| Maximum Score Words Formed by Letters| [link](pkg/leetcode/backtracking/1255.maximum_score_words_formed_by_letters.go)| Backtracking|
|1329| Sort the Matrix Diagonally| [link](pkg/leetcode/matrix/1329.sort_matrix_diagonally.go)| Matrix|
|[1457](https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/)| Pseudo-Palindromic Paths in a Binary Tree| [link](pkg/leetcode/tree/1457.ppaths_binary_tree.go)| DFS|
|1588| Sum of All Odd Length Subarrays| [link](pkg/leetcode/array/1588.sum_all_odd_length_subarrays.go)| Prefix Sum|
|1593| Split a String Into the Max Number of Unique Substrings| [link](pkg/leetcode/backtracking/1593.split_string_into_max_number_of_unique_substrings.go)| Backtracking |
|1601| Maximum Number of Achievable Transfer Requests| [link](pkg/leetcode/backtracking/1601.max_num_achievable_transf_req.go)| Backtracking |
|1647| Minimum Deletions to Make Character Frequencies Unique| [link](pkg/leetcode/greedy/1647.min_del_make_char_freq_unique.go)| Greedy |
|1723| Find Minimum Time to Finish All Jobs| [link](pkg/leetcode/backtracking/1723.find_min_time_to_finish_all_jobs.go)| Backtracking|
|1732| Find the Highest Altitude| [link](pkg/leetcode/array/1732.highest_altitude.go)| Array|
|[1770](https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations)| Maximum Score from Performing Multiplication Operations| [link](pkg/leetcode/dynamicprogramming/1770.max_score_from_mult.go)| DP|
|1799| Maximize Score After N Operations| [link](pkg/leetcode/backtracking/1799.maximize_score_after_n_operations.go)| DFS+Memo|
|[1996](https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/)| The Number of Weak Characters in the Game| [link](pkg/leetcode/greedy/1996.num_of_weak_chars.go)| Greedy|
|2014| Longest Subsequence Repeated k Times| [link](pkg/leetcode/backtracking/2014.longest_subsequence_repeated_k_times.go)| Backtracking|
|2056| Number of Valid Move Combinations On Chessboard| [link](pkg/leetcode/backtracking/2014.longest_subsequence_repeated_k_times.go)| Backtracking|
|2065| Maximum Path Quality of a Graph| [link](pkg/leetcode/backtracking/2065.max_path_quality_graph.go)| Backtracking|
|2305| Fair Distribution of Cookies| [link](pkg/leetcode/backtracking/2305.fair_distribution_cookies.go)| Backtracking|
|2364| Count Number of Bad Pairs| [link](pkg/leetcode/hashmap/2364.count_number_of_bad_pairs.go)| HashMap|
|2395| Find Subarrays With Equal Sum| [link](pkg/leetcode/contest/biweekly86.go)| Set |
|2396| Strictly Palindromic Number| [link](pkg/leetcode/contest/biweekly86.go)| Math |
|2397| Maximum Rows Covered by Columns| [link](pkg/leetcode/contest/biweekly86.go)| Backtracking |

## Daily Coding Problem

### Google problems

#### [Problem 1. Sum of two](https://github.com/serhii-soboliev/golc/blob/main/pkg/dailycodingproblems/problems/Google/problems.md#problem-1)

#### [Problem 3. Serialize and Deserialize Binary Tree](https://github.com/serhii-soboliev/golc/blob/main/pkg/dailycodingproblems/problems/Google/problems.md#problem-3)

#### [Problem 6. XOR Linked List](https://github.com/serhii-soboliev/golc/blob/main/pkg/dailycodingproblems/problems/Google/problems.md#problem-6)

### Uber problems

#### [Problem 2. Product of Array Except Self](https://github.com/serhii-soboliev/golc/blob/main/pkg/dailycodingproblems/problems/Uber/problems.md#problem-2)

### Stripe problems

#### [Problem 4. First Missing Positive](https://github.com/serhii-soboliev/golc/blob/main/pkg/dailycodingproblems/problems/Stripe/problems.md#problem-4)
