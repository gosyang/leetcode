
public class Remove_Dup_form_SA {
	public int removeDuplicates(int[] A) {
		if (A.length == 0) return 0;
		int n = 0;
		for (int i = 1; i < A.length; i++) {
			if (A[i] != A[n])
				A[++n] = A[i];
		}
		return n + 1;
	}
}
