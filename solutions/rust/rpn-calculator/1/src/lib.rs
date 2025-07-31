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
            CalculatorInput::Add => {
                if stack.len() < 2 {
                    return None;
                }
                let (right, left) = (stack.pop().unwrap(), stack.pop().unwrap());
                stack.push(left + right);
            },
            CalculatorInput::Subtract => {
                if stack.len() < 2 {
                    return None;
                }
                let (right, left) = (stack.pop().unwrap(), stack.pop().unwrap());
                stack.push(left - right);
            },
            CalculatorInput::Multiply => {
                if stack.len() < 2 {
                    return None;
                }
                let (right, left) = (stack.pop().unwrap(), stack.pop().unwrap());
                stack.push(left * right);
            },
            CalculatorInput::Divide => {
                if stack.len() < 2 {
                    return None;
                }
                let (right, left) = (stack.pop().unwrap(), stack.pop().unwrap());
                stack.push(left / right);
            },
            CalculatorInput::Value(n) => stack.push(*n),
        };
    };

    if stack.len() > 1 {
        return None;
    }

    stack.pop()
}
