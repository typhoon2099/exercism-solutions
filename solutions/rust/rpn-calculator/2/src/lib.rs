#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut stack: Vec<i32> = vec![];

    for input in inputs.iter() {
        match input {
            CalculatorInput::Value(n) => stack.push(*n),
            _ => {
                if stack.len() < 2 {
                    return None;
                }
                let (right, left) = (stack.pop(), stack.pop());

                match (left, right) {
                    (Some(left), Some(right)) => {
                        match input {
                            CalculatorInput::Divide => stack.push(left / right),
                            CalculatorInput::Multiply => stack.push(left * right),
                            CalculatorInput::Add => stack.push(left + right),
                            CalculatorInput::Subtract => stack.push(left - right),
                            _ => {},
                        }
                    },
                    _ => return None,
                }
            },
        };
    };

    if stack.len() > 1 { return None }

    stack.pop()
}
