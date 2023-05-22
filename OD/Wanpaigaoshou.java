import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Wanpaigaoshou {
    public static void main(String[] args) {

        //1,-5,-6,4,3,6,-2
        //11
        Scanner sc = new Scanner(System.in);
        String str = sc.nextLine();
        String[] split = str.split(",");
        List<Integer> list = new ArrayList<>();
        List<Integer> list_total = new ArrayList<>();
        for (int i = 0; i < split.length; i++) {
            int i1 = Integer.parseInt(split[i]);
            list.add(i1);
        }
        int total = 0;
        int total_y = 0; //选择负数牌面时，总的值
        int total_n = 0; //不负数牌面时，总的值
        for (int i = 0; i < list.size(); i++) {
            //正数的时候，必选择
            if (list.get(i)>=0){
                total+=list.get(i);
                list_total.add(total);
                //负数时，进行判断
            }else if(list.get(i)<0){
                //需要该负数时，总的值
                total_y = total+list.get(i);
                //不需要该负数时，总的值
                if(i-3>=0){
                    total_n = list_total.get(i-3);
                }else {
                    total_n = 0;
                }
                //判断需要和不需要负数时，哪个大就选择哪个
                total = Math.max(total_y,total_n);
                list_total.add(total);
            }
        }
        System.out.println(total);
    }
}
