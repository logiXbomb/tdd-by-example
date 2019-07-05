struct Money {
    amount: i32,
    currency: &'static str,
}

const USD: &'static str = "USD";
const CHF: &'static str = "CHF";

impl Money {
    fn equals(&self, m: Money) -> bool {
        self.currency == m.currency && self.amount == m.amount
    }
    fn new(amount: i32, currency: &'static str) -> Money {
        Money { amount, currency }
    }
    fn times(&self, multiplier: i32) -> Money {
        Money::new(self.amount * multiplier, self.currency)
    }
    fn plus(&self, add: Money) -> Money {
        Money::new(self.amount + add.amount, self.currency)
    }
}

#[test]
fn test_multiplication() {
    let five = Money::new(5, USD);
    assert!(Money::new(10, USD).equals(five.times(2)));
    assert!(Money::new(15, USD).equals(five.times(3)));
}

#[test]
fn test_simple_addition() {
    let sum = Money::new(5, USD).plus(Money::new(5, USD));
    assert!(Money::new(10, USD).equals(sum))
}

#[test]
fn test_multiplication_franc() {
    let five = Money::new(5, CHF);
    assert!(Money::new(10, CHF).equals(five.times(2)));
    assert!(Money::new(15, CHF).equals(five.times(3)));
}

#[test]
fn test_equality() {
    assert!(Money::new(5, USD).equals(Money::new(5, USD)));
    assert!(!Money::new(5, USD).equals(Money::new(6, USD)));
}

#[test]
fn test_equality_franc() {
    assert!(Money::new(5, CHF).equals(Money::new(5, CHF)));
    assert!(!Money::new(5, CHF).equals(Money::new(6, CHF)));
}

#[test]
fn test_different_currency_equality() {
    assert!(!Money::new(5, CHF).equals(Money::new(5, USD)))
}
