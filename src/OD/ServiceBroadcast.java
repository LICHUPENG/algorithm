package OD;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class ServiceBroadcast {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        while (in.hasNext()) {
            String[] str = in.nextLine().split(" ");
            List<Integer> list = new ArrayList<>();
            for (int i = 0; i < str.length; i++) {  // 把第一行加入list
                list.add(Integer.parseInt(str[i]));
            }
            for (int i = 0; i < str.length - 1; i++) {  // 把剩下的行都加入list
                String[] s = in.nextLine().split(" ");
                for (int j = 0; j < s.length; j++) {
                    list.add(Integer.parseInt(s[j]));
                }
            }
            int N = str.length;
            List<List<Integer>> res = new ArrayList<>(); // 存储需要广播的服务器
            for (int i = 0; i < N; i++) {  // 每一行
                if (isContainer(res, i)) {  // 判断某个服务器是否已经存在其连通的服务器集合中
                    continue;
                }
                List<Integer> newList = new ArrayList<>();
                newList.add(i);
                int lastLength = 0;
                while (lastLength != newList.size()) { // 判断当前集合可以联通的服务器
                    for (int k = 0; k < newList.size(); k++) {
                        int x = newList.get(k);
                        for (int j = 0; j < N; j++) {
                            int index = x * (N) + j;  // 找到在对应list的索引
                            if (list.get(index).equals(1)) {
                                if (!newList.contains(j)) {
                                    newList.add(j);
                                }
                            }
                        }
                    }
                    lastLength = newList.size();
                }
                res.add(newList);
            }
            System.out.println(res.size());
        }
    }
    public static Boolean isContainer(List<List<Integer>> res, int x) {
        for (List<Integer> integers : res) {
            if (integers.contains(x)) {
                return true;
            }
        }
        return false;
    }
}