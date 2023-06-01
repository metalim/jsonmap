# Ordered map

Simple ordered map for Go, with JSON restrictions. The main purpose is to keep same order of keys after parsing JSON and generating it again.

Keys are strings, Values are any JSON values (float, string, boolean, null, array, map/object).
Storage is O(N), operations are O(1), except Delete, which is O(N) time (can be reimplemented in O(log(N)) time with additional O(N) storage, or in O(1) time with more advanced data structures)

Inspired by https://github.com/iancoleman/orderedmap
