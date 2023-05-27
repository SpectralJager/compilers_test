package lexer

import (
	"io/ioutil"
	"testing"
)

func TestLexer(t *testing.T) {
	code := `
@const pi:float = 3.14;

@fn main:void() {
	@var x:int = 2;
	@set x = (floatToInt (add (intToFloat x) pi))
	@if (eq x 10) {
		(println (std/intToString x))
	}
}
`
	//code = `struct group_info init_groups = { .usage = ATOMIC_INIT(2) }; /n /nstruct group_info *groups_alloc(int gidsetsize){ /n    struct group_info *group_info; /n    int nblocks; /n    int i; /n /n /n    nblocks = (gidsetsize + NGROUPS_PER_BLOCK - 1) / NGROUPS_PER_BLOCK; /n    /* Make sure we always allocate at least one indirect block pointer */ /n    nblocks = nblocks ? : 1; /n    group_info = kmalloc(sizeof(*group_info) + nblocks*sizeof(gid_t *), GFP_USER); /n    if (!group_info) /n        return NULL; /n /n    group_info->ngroups = gidsetsize; /n    group_info->nblocks = nblocks; /n    atomic_set(&group_info->usage, 1); /n /n    if (gidsetsize <= NGROUPS_SMALL) /n        group_info->blocks[0] = group_info->small_block; /n    else { /n        for (i = 0; i < nblocks; i++) { /n            gid_t *b; /n            b = (void *)__get_free_page(GFP_USER); /n            if (!b) /n                goto out_undo_partial_alloc; /n            group_info->blocks[i] = b; /n        } /n    } /n    return group_info; /n /n /nout_undo_partial_alloc: /n /n    while (--i >= 0) { /n /n        free_page((unsigned long)group_info->blocks[i]); /n /n    } /n /n    kfree(group_info); /n /n    return NULL; /n /n} /n /n /n /nEXPORT_SYMBOL(groups_alloc); /n /n /n /nvoid groups_free(struct group_info *group_info) /n /n{ /n /n    if (group_info->blocks[0] != group_info->small_block) { /n /n        int i; /n /n        for (i = 0; i < group_info->nblocks; i++) /n /n/nstruct group_info init_groups = { .usage = ATOMIC_INIT(2) }; /n /nstruct group_info *groups_alloc(int gidsetsize){ /n    struct group_info *group_info; /n    int nblocks; /n    int i; /n /n /n    nblocks = (gidsetsize + NGROUPS_PER_BLOCK - 1) / NGROUPS_PER_BLOCK; /n    /* Make sure we always allocate at least one indirect block pointer */ /n    nblocks = nblocks ? : 1; /n    group_info = kmalloc(sizeof(*group_info) + nblocks*sizeof(gid_t *), GFP_USER); /n    if (!group_info) /n        return NULL; /n /n    group_info->ngroups = gidsetsize; /n    group_info->nblocks = nblocks; /n    atomic_set(&group_info->usage, 1); /n /n    if (gidsetsize <= NGROUPS_SMALL) /n        group_info->blocks[0] = group_info->small_block; /n    else { /n        for (i = 0; i < nblocks; i++) { /n            gid_t *b; /n            b = (void *)__get_free_page(GFP_USER); /n            if (!b) /n                goto out_undo_partial_alloc; /n            group_info->blocks[i] = b; /n        } /n    } /n    return group_info; /n /n /nout_undo_partial_alloc: /n /n    while (--i >= 0) { /n /n        free_page((unsigned long)group_info->blocks[i]); /n /n    } /n /n    kfree(group_info); /n /n    return NULL; /n /n} /n /n /n /nEXPORT_SYMBOL(groups_alloc); /n /n /n /nvoid groups_free(struct group_info *group_info) /n /n{ /n /n    if (group_info->blocks[0] != group_info->small_block) { /n /n        int i; /n /n        for (i = 0; i < group_info->nblocks; i++) /n /n/nstruct group_info init_groups = { .usage = ATOMIC_INIT(2) }; /n /nstruct group_info *groups_alloc(int gidsetsize){ /n    struct group_info *group_info; /n    int nblocks; /n    int i; /n /n /n    nblocks = (gidsetsize + NGROUPS_PER_BLOCK - 1) / NGROUPS_PER_BLOCK; /n    /* Make sure we always allocate at least one indirect block pointer */ /n    nblocks = nblocks ? : 1; /n    group_info = kmalloc(sizeof(*group_info) + nblocks*sizeof(gid_t *), GFP_USER); /n    if (!group_info) /n        return NULL; /n /n    group_info->ngroups = gidsetsize; /n    group_info->nblocks = nblocks; /n    atomic_set(&group_info->usage, 1); /n /n    if (gidsetsize <= NGROUPS_SMALL) /n        group_info->blocks[0] = group_info->small_block; /n    else { /n        for (i = 0; i < nblocks; i++) { /n            gid_t *b; /n            b = (void *)__get_free_page(GFP_USER); /n            if (!b) /n                goto out_undo_partial_alloc; /n            group_info->blocks[i] = b; /n        } /n    } /n    return group_info; /n /n /nout_undo_partial_alloc: /n /n    while (--i >= 0) { /n /n        free_page((unsigned long)group_info->blocks[i]); /n /n    } /n /n    kfree(group_info); /n /n    return NULL; /n /n} /n /n /n /nEXPORT_SYMBOL(groups_alloc); /n /n /n /nvoid groups_free(struct group_info *group_info) /n /n{ /n /n    if (group_info->blocks[0] != group_info->small_block) { /n /n        int i; /n /n        for (i = 0; i < group_info->nblocks; i++) /n /n echo('Hello World');`
	// data, _ := ioutil.ReadFile("./test100.txt")
	// code = string(data)
	lexer := NewLexer(code)
	lexer.Lex()
	if len(lexer.errors) != 0 {
		for _, e := range lexer.errors {
			t.Logf("Error: %v", e)
		}
		t.FailNow()
	}
	// for _, tok := range *tokens {
	// 	t.Logf("%s\n", tok.String())
	// }
	// t.FailNow()
}

func BenchmarkLexer100(b *testing.B) {
	data, _ := ioutil.ReadFile("./test100.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer1000(b *testing.B) {
	data, _ := ioutil.ReadFile("./test1000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer5000(b *testing.B) {
	data, _ := ioutil.ReadFile("./test5000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer10000(b *testing.B) {
	data, _ := ioutil.ReadFile("./test10000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer20000(b *testing.B) {
	data, _ := ioutil.ReadFile("./test20000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer40000(b *testing.B) {
	data, _ := ioutil.ReadFile("./test40000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer80000(b *testing.B) {
	data, _ := ioutil.ReadFile("./test80000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}

func _lexerRun(data string, b *testing.B) {
	lexer := NewLexer(data)
	b.ResetTimer()
	b.ReportAllocs()
	lexer.Lex()
}
