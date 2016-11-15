#include <vector>
#include <iostream>
#include <algorithm>
#include <cassert>
#include <string>
#include <sstream>

int pow2(int n) {
  return 1 << n;
}

char get_base64_char(int n) {
  assert(n >= 0 && n <= 63);
  if (n < 26) {
    return 65 + n;
  } else if (n < 52) {
    return n - 26 + 97;
  } else if (n < 62) {
    return n - 52 + 48;
  } else if (n == 62) {
    return '+';
  }
  return '/';
}

std::string hex2base64(std::vector<char> hexstr) {
  unsigned int bits = hexstr.size() * 8;
  char* bin_array = new char[bits];
  int k = 0;
  for (int i = 0; i < hexstr.size(); ++i) {
    char c = hexstr[i];
    for (int j = 7; j >= 0; --j) {
	bin_array[k + j] = c & 1;
	c >>= 1;
    }
    k += 8;
  }

  std::vector<char> base64_arr;
  int i = 0;
  for (; i + 5 <= bits; i = i + 6) {
    int base64_num = 0;
    for (int j = 5; j >= 0; --j) {
      base64_num += pow2(j) * bin_array[i + 5 - j];
    }
    base64_arr.push_back(get_base64_char(base64_num));
  }
  if (bits - i == 2) {
    int base64_num = pow2(5) * bin_array[i] + pow2(4) * bin_array[i + 1];
    base64_arr.push_back(get_base64_char(base64_num));
    base64_arr.push_back('=');
    base64_arr.push_back('=');
  } else if (bits - i == 4) {
    int base64_num = pow2(5) * bin_array[i] + pow2(4) * bin_array[i + 1]
      + pow2(3) * bin_array[i + 2] + pow2(2) * bin_array[i + 3];
    base64_arr.push_back(get_base64_char(base64_num));
    base64_arr.push_back('=');
  }
  delete [] bin_array;
  return std::string(base64_arr.begin(), base64_arr.end());
}


std::vector<char> str2hexstr(const std::string& str) {
  std::vector<char> hexstr;
  for (int i = 0; i < str.size(); i = i + 2) {
    std::string substr(std::string(str, i, 2));
    std::istringstream iss(substr);
    int hexval;
    iss >> std::hex >> hexval;
    hexstr.push_back(static_cast<char>(hexval));
  }
  std::for_each(hexstr.begin(), hexstr.end(),
	       [](char c) { std::cout << std::hex << c; });
  std::cout << '\n';
  return hexstr;
}

int main() {
  // std::vector<char> hexstr = {0x4d};
  std::vector<char> hexstr = str2hexstr("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d");
  std::string base64 = hex2base64(hexstr);
  assert(base64 == "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t");
}
