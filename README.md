# Serif Health Anthem PPO URL Extraction Project

## Project Overview

This project aims to extract URLs for Anthem PPO plans in New York state from a large JSON index file. The goal is to identify and isolate the relevant URLs from a complex and extensive dataset.

### Initial Approach

- Used a combination of URL and description to identify relevant files.
- Implemented a simple filtering mechanism based on keywords "PPO" and "New York" in the description.
- Utilized Go's built-in JSON decoding and gzip decompression to handle the large input file.

### Observations and Challenges

- The initial pass yielded about 56 files, indicating potential over-inclusion.
- URLs contained subdomains for insurances not related to New York (e.g., anthembcbsco for Colorado).
- The description field proved unreliable for accurate identification of New York PPO plans.

### Current Approach

The current implementation uses two main criteria to identify New York PPO plans:

1. The subdomain "empirebcbs" to identify New York plans.
2. The presence of "ppo" in the description to identify PPO plans.

### Identified Challenges with Current Approach

- **Plan Identifier Variability**: The plan identifiers in the URLs are more varied than initially thought (e.g., "301_71A0", "800_72A0", "302_42B0", "020_02I0", "361_50I0"). This makes it difficult to use them as reliable indicators of PPO plans.
- **Reliance on Description**: Using the description to identify PPO plans may not be entirely reliable, as it could lead to both false positives and false negatives.
- **Lack of Standardized** Identification: There doesn't seem to be a standardized way to identify PPO plans solely from the URL structure.
- **Connecting Reporting Plans**: Difficulty in connecting reporting plans with file locations to better identify and uniquely identify plans.
- **Multi-part Files**: Implementing a completeness check for multi-part files.

### Potential Improvements

- **EIN Utilization**: Investigate the use of Employer Identification Numbers (EINs) or other unique identifiers for more precise plan identification.
- **Deeper Analysis**: Conduct a thorough analysis of the relationship between plan identifiers and plan types to potentially derive a pattern or rule set.
- **Additional Data Sources**: Explore the possibility of using additional data from the index file or external sources to improve plan type identification.
- **Expert Consultation**: Consider reaching out to Anthem or industry experts for insights into their naming conventions and plan identification methods.
- **Robust Validation**: Implement a more robust validation system for multi-part files.

### Next Steps

- Conduct further research on plan identifiers and their relationship to plan types.
- Investigate the structure of the Reporting Plans Object and how it can be utilized for more accurate plan identification.
- Analyze the full dataset to identify any patterns or consistencies that could improve the identification process.
- Consider developing a more sophisticated algorithm that takes into account multiple factors (subdomain, plan identifier, description, EIN) to classify plans.

### Conclusion

While progress has been made in isolating New York PPO plan URLs, significant challenges remain in achieving high accuracy and completeness. The project demonstrates the complexity of working with large, inconsistently structured datasets and the importance of iterative refinement in data processing tasks. Further research and possibly additional data sources will be necessary to refine the identification process and improve the accuracy of the results.

### Timebox Constraint

Due to the need to understand the context related to insurance and the likes, the solution took over 2 hours to research and gain a good understanding of terms like MRF and PPO. Once that was determined, a better gauge of what needed to be done was achieved, allowing for more efficient time-boxing. The total time spent on coding the solution was limited to approximately 2 hours as per the requirements.

## File Descriptions

- `first_pass_through.txt`: This file contains intermediate data from the first pass through the JSON index file. It includes all potential matches before applying more stringent filtering criteria.
- `second_pass_through.txt`: This file contains data from the second pass through the JSON index file. It includes a refined list of matches after applying additional filtering criteria and improvements to the extraction logic.
- `chunk.json`: A sample chunk of the JSON data used for testing and validation purposes. This file helps in understanding the structure and content of the index file, allowing for better development and debugging of the extraction logic.

## Instructions to Run the Script

1. Dependencies: Ensure you have Go installed on your machine.
2. Input File: Place the `anthem_Index_2024-07-01.json.gz` file in the root directory of the project. If you decide to name it differently, update the main.go inputfile variable accordingly.
3. Run the Script: Execute the script using the command `go run main.go`.
4. Output: The resulting URLs will be written to `anthem_ny_ppo_urls.txt`.

### Example Command

```sh
go run main.go
```

This README now provides a comprehensive overview of the project, detailing the initial and current approaches, observed challenges, potential improvements, next steps, and instructions for running the script.

### Time Taken to Run the Code

The script processes a large JSON file, which may take a few minutes depending on the size of the file and the performance of the machine. Typically, it should take less than 10 minutes to run the script on a standard machine.

### Tradeoffs Made

Efficiency vs. Simplicity: Chose to handle JSON parsing incrementally to avoid memory issues while keeping the code simple.
Pattern Matching: Focused on specific patterns in descriptions for efficient filtering, knowing it might not capture all edge cases.
Time Constraints: Limited the depth of validation and robustness of error handling due to the 2-hour time constraint.

### Future Improvements

Implement more robust error handling and logging.
Optimize performance by parallelizing URL extraction.
Use additional data sources or algorithms for better plan identification.
