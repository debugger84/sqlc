# Code generated by sqlc. DO NOT EDIT.
# versions:
#   sqlc v1.23.0
import dataclasses


@dataclasses.dataclass()
class Bar:
    id: int
    name: str


@dataclasses.dataclass()
class Exclusions:
    id: int
    name: str


@dataclasses.dataclass()
class MyData:
    id: int
    name: str
