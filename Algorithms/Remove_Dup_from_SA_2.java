
public class Remove_Dup_from_SA_2 {
    public int removeDuplicates(int[] A) {
    	if (A.length < 2) return A.length;
       int n = 0;
       int time = 0;
       for (int i = 1; i < A.length; i++ ) {
    	   if (A[i] == A[n]) 
    		   time++;
    	   if (A[i] != A[n]) {
    		   A[++n] = A[i];
    		   time = 0;
    	   } else if (A[i] == A[n] && time < 2) {
    		   A[++n] = A[i];
    	   }
        }
       return n+1;
    }
}
