from matplotlib import pyplot as plot
from subprocess import run, PIPE

results = run("go test -bench .", shell=True, stdout=PIPE, stderr=PIPE)
if results.returncode:
    print("Tests didn't pass!",
          "STDOUT:",
          results.stdout,
          "STDERR:",
          results.stderr,
          "Return code:",
          results.returncode,
          sep='\n')

lengths = [1, 4, 16, 64, 256, 1024]
times = [int(s.split()[2]) / 1000
         for s
         in results.stdout.split(b'\n')
         if s.startswith(b"BenchmarkRandomString")]
alphanumeric_times = [int(s.split()[2]) / 1000
                      for s
                      in results.stdout.split(b'\n')
                      if s.startswith(b"BenchmarkRandomAlphanumericString")]

plot.plot(lengths, times)
plot.plot(lengths, alphanumeric_times)
plot.ylabel("ms/op")
plot.autoscale()
plot.show()
