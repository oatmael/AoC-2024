package aocutils

func ArrayReduce[T ~[]TElement, TElement any, Acc any](
	array T, 
	initial Acc, 
	callback func(element TElement, index int, acc Acc) Acc,
) Acc {
	var acc Acc;
	acc = initial;

	for i, element := range array {
		acc = callback(element, i, acc);
	}

	return acc;
}

func ArrayMap[T ~[]TElement, TElement any, M ~[]MElement, MElement any](
	array T,
	callback func(element TElement, index int) MElement,
) M {
	var mapped []MElement;

	for i, element := range array {
		mapped = append(mapped, callback(element, i));
	}

	return mapped;
}

func ArrayFilter[T ~[]TElement, TElement any](
	array T, 
	callback func (element TElement, index int) bool,
) T {
	var filtered []TElement;

	for i, element := range array {
		if callback(element, i) {
			filtered = append(filtered, element);
		} 
	}

	return filtered;
}