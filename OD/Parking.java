import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Parking {
    public static void main(String[] args) {

        Scanner sc = new Scanner(System.in);

        String[] s = sc.nextLine().split(",");

        List<Integer> list = new ArrayList<>();

        for(int i=0;i<s.length;i++){
            if(s[i].equals("1")){
                list.add(i);    //记录1的下标，0的长度=后一个下标 - 前一个下标 - 1
            }
        }

        int maxFE = 0;
        if(s[0].equals("0")){   //考虑到开头为0的情形
            maxFE = Math.max(maxFE,list.get(0));
        }
        if(s[s.length-1].equals("0")){  //考虑到末尾为0的情形
            maxFE = Math.max(maxFE,s.length-1-list.get(list.size()-1));
        }

        int max = 0;

        for(int i=1;i<list.size();i++){
            max = Math.max(max,list.get(i)-list.get(i-1)-1);    //通过1的下标求0的长度
        }

        if(max%2==0){
            System.out.println(Math.max(maxFE,max/2));  //偶数直接除以2
        }else {
            System.out.println(Math.max(maxFE,(max+1)/2));  //奇数+1除以2
        }
    }
}
