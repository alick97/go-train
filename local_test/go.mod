module example.com/ee

go 1.21.3

replace example.com/p1 => ./p1

replace example.com/m1 => ./m1

require example.com/m1 v0.0.0-00010101000000-000000000000
