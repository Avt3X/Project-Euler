import java.util.ArrayList;
import java.util.HashSet;
import java.util.function.Supplier;

public class Problem23 {

    public static void main(String[] args) {
        Supplier<Integer> solver = () -> {
            int result = 0;
            var others = new HashSet<Integer>();
            var abundands = new ArrayList<Integer>();
            for (int i = 1; i <= 28123; i++) {
                if (divSum(i) > i) {
                    abundands.add(i);
                }
            }

            for (int i = 0; i < abundands.size(); i++) {
                for (int i1 = i; i1 < abundands.size(); i1++) {
                    others.add(abundands.get(i) + abundands.get(i1));
                }
            }

            for (int i = 1; i <= 28123; i++) {
                if (!others.contains(i)) {
                    result += i;
                }
            }

            return result;
        };

        System.out.println(time(solver));
    }

    private static int divSum(int num) {
        int sum = 1;
        int sqrt = ((int) Math.sqrt(num));
        for (int i = 2; i <= sqrt; i++) {
            if (i * i == num) {
                sum += i;
                continue;
            }
            if (num % i == 0) {
                sum += i;
                sum += num / i;
            }
        }

        return sum;
    }

    private static long time(Supplier<Integer> solver) {
        long start = System.currentTimeMillis();

        solver.get();

        return System.currentTimeMillis() - start;
    }
}
