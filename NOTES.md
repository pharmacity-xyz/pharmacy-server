# NOTES

# MODULE 1: Introduction to amazon web services

## What is AWS?

With cloud computing, companies do not have to manage and maintain their own hardware and data centers. Instead, companies like AWS own and maintain data centers and provide virtual data center technologies and services to companies and users over the internet.

AWS provides cloud computing services.

## AWS Global Infrastructure

Infrastructure, like data centers and networking connectivity, still exists as the foundation of every cloud application.

## Interacting with AWS

1. The AWS Management Console

2. The AWS Command Line Interface (AWS CLI)

3. AWS SDKs
   Developers commonly use AWS SDKs to integrate their application source code with AWS services.

## Security and the AWS Shared Responsibility Model

### AWS responsibility

AWS is responsible for security of the cloud.

- Protecting and securing AWS Regions, Availability Zones, and data centers, down to the phisical security of the buildings
- Managing the hardware, software, and networking components that run AWS services, such as the phisical servers, host operating systems, virtualization layers, and AWS networking components

### Customer responsibility

Customers are responsible for security in the cloud.

## Protect the AWS Root User

## AWS Identity and Access Management

### IAM

AWS Identity and Access Management (IAM) is an AWS service that helps you manage access to your AWS account and resources. It also provides a centralized view of who and what are allowed inside your AWS account (authentication), and who and what have permissions to use and work with your AWS resources (authorization).

## Role-Based Access in AWS

# MODULE 2: AWS Compute

## Compute as a Service

The first building block you need to host an application is a server. Servers usually can handle Hypertext Transfer Protocol (HTTP) requests and send responses to clients following the clinet-server model, although any API-based communication also falls under this model.

## Amazon Elastic Compute Cloud

### Amazon EC2

Amazon EC2 is a web service that provides secure, resizable compute capacity in the cloud.

You can create and manage EC2 instances through the AWS Management Console, the AWS Command Line Interface (CLI), AWS software development kits (SDKs), automation tools, and infrastructure orchestration services.

### Amazon Machine Image

In the AWS Cloud, the OS installation is not your responsibility. Instead, it's built into the AMI that you choose.

### Relationship between AMIs and EC2 instances

EC2 instances are live instantiations of what is defined in an AMI, much like a cake is a live instantiation of a cake recipe. you can also see this kind of relationship between a class and an object

## Amazon EC2 Instance Lifecycle

## Container Services

### Containers

Docker - is a pupular container runtime that simplifies the management of the entire operating system stack needed for container isolation, including networking and storage.

### Difference between containers and virtual machines (VMs)

Containers share the same operating system and kernel as the host they exist on, whereas virtual machines contain their own operating system.

### Orchestrate containers

In AWS, containers run on EC2 instances.

### Use Kubernetes with Amazon Elastic Kubernetes Service (Amazon EKS)

Kubernetes is a portable, extensible, open source platform for managing containerized workloads and services.

## Serverless

## AWS Lambda

# Module 3: AWS NETWORKING

## Networking

## Amazon Virtual Private Cloud

A virtual private cloud (VPC) is an isolated network that you create in the AWS Cloud, similar to a traditional network in a data center.

# Module 4: AWS Storage

## Amazon EC2 instance Storage and Amazon Elastic Block Store

### Amazon EC2 instance store

Amazon EC2 instance store provides temporary block-level storage for an instance. This storage is located on disks that are physically attached to the host computer.

### Amazon Elastic Block Storage (Amazon EBS)

As the name implies, Amazon EBS is a block-level storage devide that you can attach to an Amazon EC2 instance.

## Object Storage with Amazon Simple Storage Service

Amazon S3 is an object storage service. Object storage stores data in a flat structure, using unique identifiers to look up objects when requested.

### Amazon S3 use cases

- backup and storage
- media hosting
- software delivery
- data lakes
- static websites
- static content

# Module 5: Databases

## Databases on AWS

### Relational databases

A table stores data in rows and columns. A row, often called a record, contains all information about a specific entry. Columns describe attributes of an entry.

## Amazon Relational Database Service

### Amazon RDS

Amazon Relational Database Service (Amazon RDS) lets customers create and manage relational databases in the cloud without the operational burden of traditional database management.

## Purpose-Built Databases

1. Amazon Relational Databases(RDS) - Default
2. Amazon Dynamo DB(DynamoDB) - Non-relational DB (key - value store)
3. Amazon Document DB(Amazon DocumentDB) - MongoDB compatibility
4. Amazon Neptune graph database(Neptune) - fraud detection
5. Amazon QLDB ledger database(Amazon QLDB) - immutability

## Amazon DynamoDB

### Amazon DynamoDB introduction

Amazon DynamoDB is a fully managed NoSQL database service that provides fast and predictable performance with seamless scalability. 

### Amazon DynamoDB core components
In DynamoDB, tables, items, and attributes are the core components that you work with. A table is a collection of items, and each item is a collection of attributes. 

- **Tables**: Similar to other database systems, DynamoDB stores data in tables. A table is a collection of data. 
- **Items**: Each table contains zero or more items. An item is a group of attributes that is uniquely identifiable among all the other items. 
- **Attributes**: Each item is composed of one or more attributes. An attribute is a fundamental data element, something that does not need to be broken down any further. 

## Choose the Right Database Service

![Screenshot from 2023-02-06 11-20-56](https://user-images.githubusercontent.com/62386689/216883055-4b601463-c56b-4587-8111-0e1f278c00b3.png)
