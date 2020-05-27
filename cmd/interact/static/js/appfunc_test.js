QUnit.test( "hello test", function( assert ) {
  assert.ok( 1 == "1", "Passed!" );
});

QUnit.test( "autoComplete", (t) => {
  let exp = ['a','a','b','c'];
  let v = autoComplete('a');
  t.deepEqual(v, exp, `got: ${v} expected: ${exp}`);
  t.ok(v.join(',') == exp.join(','), `got: ${v} expected: ${exp}`);
  t.equal(v.join(','), exp.join(','), `got: ${v} expected: ${exp}`);
});
