import java.util.Scanner;

public class StringSplit {
        public static void main(String[] args) {
            Scanner input = new Scanner(System.in);
            while(input.hasNext()) {
                String s = input.nextLine();
                int count = 0;	// count范围[0,8]，当count=8时立刻变成0
                StringBuilder sb = new StringBuilder();	// 拼接字符串
                for(int i = 0; i < s.length(); ++i) {
                    sb.append(s.charAt(i));
                    ++count;
                    if(i == s.length() - 1) {	// 最后不足8个字符的补零
                        for(int j = 0; j < 8 - count; ++j) {
                            sb.append("0");
                        }
                        System.out.println(sb.toString());
                        break;
                    }
                    if(count == 8) {
                        count = 0;
                        System.out.println(sb.toString());	// 每8个字符输出一遍
                        sb = new StringBuilder();	// 每8个字符重new一个对象，相当于清空
                        continue;
                    }
                }
            }
        }
}
