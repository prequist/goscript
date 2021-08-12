export function pluralize(word: string, count: i8): string {
	if (count === 1) {
		return word;
	}

	return word + "s";
}
