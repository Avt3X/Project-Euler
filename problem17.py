d = {0: '', 1: 'one', 2: 'two', 3: 'three', 4: 'four', 5: 'five', 6: 'six', 7: 'seven', 8: 'eight', 9: 'nine',
     10: 'ten',
     11: 'eleven', 12: 'twelve',
     13: 'thirteen', 14: 'fourteen', 15: 'fifteen', 16: 'sixteen', 17: 'seventeen', 18: 'eighteen', 19: 'nineteen',
     20: 'twenty', 30: 'thirty', 40: 'forty', 50: 'fifty', 60: 'sixty', 70: 'seventy', 80: 'eighty', 90: 'ninety',
     100: 'onehundred', 1000: 'onethousand'}

res = 0

for i in range(1, 1001):
    s = ''
    r = 0
    if i in d:
        r += len(d[i])
    elif i % 100 == 0:
        tmp = i
        s = d[tmp/100] + 'hundred'
        r += len(s)
        d[i] = s
    else:
        tmp = i
        if tmp > 99:
            j = 2
            r += 3
            s += 'and'
            d[i] = s
        else:
            j = 1
        while tmp != 0:
            r += len(d[tmp % (10 ** j)])
            s += d[tmp % (10 ** j)]
            tmp -= tmp % (10 ** j)
            j += 1
        d[i] = s
    res += r

print(res)
