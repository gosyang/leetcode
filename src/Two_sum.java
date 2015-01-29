import java.util.Hashtable;
public class Two_sum {
	public int[] twoSum(int[] numbers, int target) {
		int[] a = {0,0};
		Hashtable<Integer, Integer> hash = new Hashtable<Integer, Integer>();
		for (int i = 0; i <= numbers.length-1; i++) {
			hash.put(numbers[i], i);
		}
		for (int i = 0; i <= numbers.length-1; i++) {
			Integer n = hash.get(target - numbers[i]);
			if (n != null && n!=i) {
				if (n<i) {
					a[0] = n + 1;
					a[1] = i + 1;
					return a;
				} else {
					a[0] = i + 1;
					a[1] = n + 1;
					return a;
				}
			}
		}
	return a;
	}
}
