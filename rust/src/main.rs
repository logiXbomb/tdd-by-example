struct Dollar {
    amount: i32,
}

impl Dollar {
    fn new(amount: i32) -> Dollar {
        Dollar { amount }
    }
    fn times(&self, multiplier: i32) -> Dollar {
        Dollar::new(self.amount * multiplier)
    }
    fn equals(&self, dollar: Dollar) -> bool {
        self.amount == dollar.amount
    }
}

#[test]
fn test_multiplication() {
    let five = Dollar::new(5);
    assert!(Dollar::new(10).equals(five.times(2)));
    assert!(Dollar::new(15).equals(five.times(3)));
}

#[test]
fn test_equality() {
    assert!(Dollar::new(5).equals(Dollar::new(5)));
    assert!(!Dollar::new(5).equals(Dollar::new(6)));
}
