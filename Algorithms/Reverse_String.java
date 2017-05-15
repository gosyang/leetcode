
public class Reverse_String {
    public String reverseWords(String s) {

        String[] ss = s.split(" ");
        String ans = "";
        for (String t : ss) {
        	if (!t.equals("")) {
            	ans = t + " " + ans;
        	}
        }
        return ans.length()==0?"":ans.substring(0,ans.length()-1);
    }
}
