package utils

func Map(vs []string, f func(string) string) []string {
  vsm := make([]string, len(vs))
  for i, v := range vs {
    vsm[i] = f(v)
  }
  return vsm
}

func Filter(vs []string, f func(string) bool) []string {
  vsf := make([]string, 0)
  for _, v := range vs {
    if f(v) {
      vsf = append(vsf, v)
   }
  }
  return vsf
}

func FilterEmptyStrings(strs []string) []string {
  without := Filter(strs, func(v string) bool {
    return !(len(v) == 0)
  })
  return without
}

