export const capitalize = function(str: string) {
	if (typeof str !== 'string' || !str) return str;
	return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
};
