import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Url {
    public static void main(String[] str) {

        Scanner sc = new Scanner(System.in);
        List<String> strs = new ArrayList();
        for (int i = 0; i < 2; i++) {
            strs.add(sc.nextLine());
        }

        String s1 = strs.get(0);
        String s2 = strs.get(1);

        String buildUrl = getUrlStr(s1) + getUrlStr(s2);
        // 去重
        String finalString = buildUrl.replace("//", "/");

        // 如果最后一个是/,去掉
        int length = finalString.length();
        if (length > 1) {
            if (finalString.endsWith("/")) {
                finalString = finalString.substring(0, length - 1);
            }
        }
        System.out.println(finalString);
    }

    /**
     * @param s
     * @return
     */
    private static String getUrlStr(String s) {
        if (s == null || s.length() == 0) {
            return "/";
        }

        String ss = s;
        char[] chars = s.toCharArray();
        if ('/' != chars[0]) {
            ss = "/" + s;
        }
        return ss;
    }
}
