package aocutils

func ArrayReduce[T ~[]TElement, TElement any, Acc any](
	array T, 
	initial Acc, 
	callback func(TElement, int, Acc) Acc,
) Acc {
	var acc Acc;
	acc = initial;

	for i, element := range array {
		acc = callback(element, i, acc);
	}

	return acc;
}

func ArrayMap[T ~[]TElement, TElement any, MElement any](
	array T,
	callback func(TElement, int) MElement,
) []MElement {
	var mapped []MElement;

	for i, element := range array {
		mapped = append(mapped, callback(element, i));
	}

	return mapped;
}

func ArrayFilter[T ~[]TElement, TElement any](
	array T, 
	callback func (TElement, int) bool,
) T {
	var filtered []TElement;

	for i, element := range array {
		if callback(element, i) {
			filtered = append(filtered, element);
		} 
	}

	return filtered;
}