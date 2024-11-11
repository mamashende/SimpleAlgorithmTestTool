import random
import string

def generate_lcs_test_cases(num_cases, str_length):
    test_cases = []
    for _ in range(num_cases):
        str1 = ''.join(random.choices(string.ascii_uppercase, k=str_length))
        str2 = ''.join(random.choices(string.ascii_uppercase, k=str_length))
        test_cases.append((str1, str2))
    return test_cases

if __name__ == "__main__":
    #num_cases = 
    str_lengths = [10, 50, 100,200,400, 500,800, 1000]
    with open("testdata.txt", "w") as f:
        for str_length in str_lengths:
            test_cases = generate_lcs_test_cases(20, str_length)
            for str1, str2 in test_cases:
                f.write(f"{str1} {str2}\n")


