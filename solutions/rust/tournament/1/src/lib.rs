use std::cmp::Ordering;
use std::collections::HashMap;

#[derive(PartialEq, Ord, Eq)]
struct Team {
    name: String,
    won: u8,
    lost: u8,
    drew: u8,
}

impl Team {
    fn played(&self) -> u8 {
        self.won + self.lost + self.drew
    }

    fn points(&self) -> u8 {
        self.won * 3 + self.drew
    }

    fn add_game(&mut self, game_result: &str) {
        match game_result {
            "win" => {
                self.won += 1;
            },
            "loss" => {
                self.lost += 1;
            },
            "draw" => {
                self.drew += 1;
            },
            _ => {},
        };
    }
}

impl PartialOrd for Team {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        if self.points() == other.points() {
            return Some(self.name.cmp(&other.name));
        }

        Some(other.points().cmp(&self.points()))
    }
}

pub fn tally(match_results: &str) -> String {
    let mut teams = HashMap::new();
    match_results.lines().for_each(|line| {
        let parts: Vec<&str> = line.split(";").collect();

        let first = teams.entry(parts[0]).or_insert(Team {
            name: parts[0].to_string(),
            won: 0,
            lost: 0,
            drew: 0,
        });

        first.add_game(parts[2]);

        let second = teams.entry(parts[1]).or_insert(Team {
            name: parts[1].to_string(),
            won: 0,
            lost: 0,
            drew: 0,
        });

        match parts[2] {
            "win" => second.add_game("loss"),
            "loss" => second.add_game("win"),
            "draw" => second.add_game("draw"),
            &_ => {},
        }
    });

    let mut output = vec![];
    output.push("Team                           | MP |  W |  D |  L |  P".to_string());

    let mut sorted_teams = teams.values().collect::<Vec<&Team>>();
    sorted_teams.sort();

    for team in sorted_teams.iter() {
        output.push(
            format!(
                "{: <30} |  {} |  {} |  {} |  {} |  {}",
                team.name,
                team.played(),
                team.won,
                team.drew,
                team.lost, 
                team.points(),
            ),
        );
    }

    output.join("\n")
}
